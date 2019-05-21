package storage

type (
	Media struct {
		ID           int64  `json:"id" db:"MediaID"`
		Name         string `json:"name" db:"MediaName"`
		Path         string `json:"hash" db:"Path"`
		Duration     int64  `json:"duration" db:"Duration"`
		LastModified int64  `json:"last_modified" db:"LastModified"`
	}

	Profile struct {
		ID   int64  `json:"id" db:"ProfileID"`
		Name string `json:"name" db:"ProfileName"`
	}

	ProfileViewingInfo struct {
		MediaID   int64 `json:"media_id" db:"MediaID"`
		ProfileID int64 `json:"profile_id" db:"ProfileID"`
		TimePoint int64 `json:"time_point" db:"TimePoint"`
		Timestamp int64 `json:"timestamp" db:"Timestamp"`
	}
)
