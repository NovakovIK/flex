package data

type Media struct {
	MediaID  int64  `json:"media_id"`
	Name     string `json:"name"`
	Hash     []byte `json:"hash"`
	Duration int64  `json:"duration"`
}

type Profile struct {
	ProfileID int64  `json:"profile_id"`
	Name      string `json:"name"`
}

type ProfileViewingInfo struct {
	MediaID   int64 `json:"media_id"`
	ProfileID int64 `json:"profile_id"`
	TimePoint int64 `json:"time_point"`
	Timestamp int64 `json:"timestamp"`
}
