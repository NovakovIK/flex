package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type (
	Storage struct {
		DB       *sqlx.DB
		MediaDAO MediaDAO
	}

	MediaDAO struct {
		DB *sqlx.DB
	}
)

func NewStorage() *Storage {
	db, err := sqlx.Connect("sqlite3", "flex.db")
	if err != nil {
		log.Fatalln(err)
	}

	storage := &Storage{
		DB:       db,
		MediaDAO: MediaDAO{db},
	}

	storage.init()

	return storage
}

func (s *Storage) init() {
	s.DB.MustExec(`
	create table if not exists media
	(
		MediaID integer
			constraint media_pk
				primary key autoincrement,
		MediaName text not null,
		Path text not null,
		Status integer not null,
		Created integer not null,
		Duration integer not null,
		LastSeen integer not null,
		TimePoint integer not null,	
		Thumbnail blob not null,
		Width integer not null,
		Height integer not null
	);
	
	create unique index if not exists media_Path_uindex
		on media (Path);
`)
}

func (d *MediaDAO) FetchAll() ([]Media, error) {
	media := make([]Media, 0)
	if err := d.DB.Select(&media, "select * from media order by LastSeen desc, Created desc"); err != nil {
		return nil, err
	}
	return media, nil
}
func (d *MediaDAO) FetchByID(id int) ([]Media, error) {
	media := make([]Media, 0)
	if err := d.DB.Select(&media, "select * from media where MediaID = $1", id); err != nil {
		return nil, err
	}
	return media, nil
}
func (d MediaDAO) Insert(media Media) error {
	_, err := d.DB.Exec(
		"insert into media(MediaName, Path, Status, Created, Duration, TimePoint, LastSeen, Thumbnail, Width, Height) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		media.Name,
		media.Path,
		media.Status,
		media.Created,
		media.Duration,
		media.TimePoint,
		media.LastSeen,
		media.Thumbnail,
		media.Width,
		media.Height,
	)
	return err
}
func (d MediaDAO) DeleteByPath(path string) error {
	_, err := d.DB.Exec("delete from media where Path = $1", path)
	return err
}

func (d *MediaDAO) Update(mediaID int, name string, lastSeen int, timePoint int) (*Media, error) {
	_, err := d.DB.Exec("update media set MediaName = $1, LastSeen = $2, TimePoint = $3 where MediaID = $4", name, lastSeen, timePoint, mediaID)

	media := &Media{}
	err = d.DB.Get(media, "select * from media where MediaID = $1", mediaID)
	if err != nil {
		return nil, err
	}
	return media, nil
}
