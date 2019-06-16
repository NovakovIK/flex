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
		ID           int         `json:"id" db:"MediaID"`
		Name         string      `json:"name" db:"MediaName"`
		Path         string      `json:"path" db:"Path"`
		Duration     int         `json:"duration" db:"Duration"`
		LastModified int         `json:"last_modified" db:"LastModified"`
		Status       MediaStatus `json:"status" db:"Status"`
	}

	Profile struct {
		ID   int    `json:"id" db:"ProfileID"`
		Name string `json:"name" db:"ProfileName"`
	}

	ProfileViewingInfo struct {
		MediaID   int `json:"media_id" db:"MediaID"`
		ProfileID int `json:"profile_id" db:"ProfileID"`
		TimePoint int `json:"time_point" db:"TimePoint"`
		Timestamp int `json:"timestamp" db:"Timestamp"`
	}
)
