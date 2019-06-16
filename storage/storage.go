package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type (
	Storage struct {
		DB                    *sqlx.DB
		MediaDAO              MediaDAO
		ProfileDAO            ProfileDAO
		ProfileViewingInfoDAO ProfileViewingInfoDAO
	}

	MediaDAO struct {
		DB *sqlx.DB
	}

	ProfileDAO struct {
		DB *sqlx.DB
	}

	ProfileViewingInfoDAO struct {
		DB *sqlx.DB
	}
)

func NewStorage() *Storage {
	db, err := sqlx.Connect("sqlite3", "flex.db")
	if err != nil {
		log.Fatalln(err)
	}

	storage := &Storage{
		DB:                    db,
		MediaDAO:              MediaDAO{db},
		ProfileDAO:            ProfileDAO{db},
		ProfileViewingInfoDAO: ProfileViewingInfoDAO{db},
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
		Duration integer not null,
		LastModified integer not null,
		Status integer not null
	);
	
	create unique index if not exists media_Path_uindex
		on media (Path);
	
	create table if not exists profile
	(
		ProfileID integer
			constraint profile_pk
				primary key autoincrement,
		ProfileName text not null
	);
	
	create unique index if not exists profile_ProfileName_uindex
		on profile (ProfileName);
	
	create table if not exists profile_viewing_info
	(
		MediaID integer
			references media
				on update cascade on delete cascade,
		ProfileID integer
			references profile
				on update cascade on delete cascade,
		TimePoint integer not null,
		Timestamp integer not null,
		constraint profile_viewing_info_pk
			primary key (MediaID, ProfileID)
	);
`)
}

func (d *MediaDAO) FetchAll() ([]Media, error) {
	media := make([]Media, 0)
	if err := d.DB.Select(&media, "select * from media order by LastModified desc"); err != nil {
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
		"insert into media(MediaName, Path, Duration, LastModified, Status) values ($1, $2, $3, $4, $5)",
		media.Name, media.Path, media.Duration, media.LastModified, media.Status,
	)
	return err
}
func (d MediaDAO) DeleteByPath(path string) error {
	_, err := d.DB.Exec("delete from media where Path = $1", path)
	return err
}

func (d *ProfileDAO) FetchAll() ([]Profile, error) {
	profiles := make([]Profile, 0)
	if err := d.DB.Select(&profiles, "select * from profile order by ProfileID asc"); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (d *ProfileDAO) FetchByID(id int) ([]Profile, error) {
	profiles := make([]Profile, 0)
	if err := d.DB.Select(&profiles, "select * from profile where ProfileID = $1", id); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (d *ProfileDAO) New(name string) (*Profile, error) {
	if _, err := d.DB.Exec("insert into profile(ProfileName) VALUES ($1)", name); err != nil {
		return nil, err
	}

	profile := &Profile{}
	if err := d.DB.Get(profile, "select * from profile where ProfileName = $1", name); err != nil {
		return nil, err
	}
	return profile, nil
}

func (d *ProfileDAO) Update(id int, name string) (*Profile, error) {
	if _, err := d.DB.Exec("update profile set ProfileName = $1 where ProfileID = $2", name, id); err != nil {
		return nil, err
	}

	profile := &Profile{}
	if err := d.DB.Get(profile, "select * from profile where ProfileName = $1", name); err != nil {
		return nil, err
	}
	return profile, nil
}

func (d *ProfileViewingInfoDAO) FetchByProfileID(id int) ([]ProfileViewingInfo, error) {
	viewingInfo := make([]ProfileViewingInfo, 0)
	if err := d.DB.Select(&viewingInfo, "select * from profile_viewing_info where ProfileID = $1 order by Timestamp desc", id); err != nil {
		return nil, err
	}
	return viewingInfo, nil
}

func (d *ProfileViewingInfoDAO) FetchByMediaID(id int) ([]ProfileViewingInfo, error) {
	viewingInfo := make([]ProfileViewingInfo, 0)
	if err := d.DB.Select(&viewingInfo, "select * from profile_viewing_info where MediaID = $1 order by Timestamp desc", id); err != nil {
		return nil, err
	}
	return viewingInfo, nil
}

func (d *ProfileViewingInfoDAO) FetchByMediaIDAndProfileID(mediaID, profileID int) ([]ProfileViewingInfo, error) {
	viewingInfo := make([]ProfileViewingInfo, 0)
	if err := d.DB.Select(&viewingInfo, "select * from profile_viewing_info where MediaID = $1 and ProfileID = $2", mediaID, profileID); err != nil {
		return nil, err
	}
	return viewingInfo, nil
}

func (d *ProfileViewingInfoDAO) FetchAll() ([]ProfileViewingInfo, error) {
	viewingInfo := make([]ProfileViewingInfo, 0)
	if err := d.DB.Select(&viewingInfo, "select * from profile_viewing_info order by Timestamp desc"); err != nil {
		return nil, err
	}
	return viewingInfo, nil
}

func (d *ProfileViewingInfoDAO) UpdateOrInsert(mediaID, profileID, timePoint, timestamp int) (*ProfileViewingInfo, error) {
	if _, err := d.DB.Exec(`
insert into profile_viewing_info(MediaID, ProfileID, TimePoint, Timestamp) values ($1, $2, $3, $4)
on conflict(MediaID, ProfileID) do update set TimePoint = $3, Timestamp = $4`, mediaID, profileID, timePoint, timestamp); err != nil {
		return nil, err
	}

	info := &ProfileViewingInfo{}
	if err := d.DB.Get(info, "select * from profile_viewing_info where MediaID = $1 and ProfileID = $2", mediaID, profileID); err != nil {
		return nil, err
	}
	return info, nil
}
