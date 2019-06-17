package storage

const (
	Encoding MediaStatus = iota + 1
	Available
)

func (s MediaStatus) String() string {
	switch s {
	case Encoding:
		return "Encoding"
	case Available:
		return "Available"
	}
	panic("Logic error")
}

type (
	MediaStatus int
	Media       struct {
		ID        int         `json:"id" db:"MediaID"`
		Name      string      `json:"name" db:"MediaName"`
		Path      string      `json:"path" db:"Path"`
		Status    MediaStatus `json:"status" db:"Status"`
		Created   int         `json:"created" db:"Created"`
		Duration  int         `json:"duration" db:"Duration"`
		LastSeen  int         `json:"last_seen" db:"LastSeen"`
		TimePoint int         `json:"time_point" db:"TimePoint"`
		Thumbnail []byte      `json:"thumbnail" db:"Thumbnail"`
	}
)
