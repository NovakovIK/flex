package flex

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Media struct {
	gorm.Model
	MediaID  int64  `json:"media_id"`
	Name     string `json:"name"`
	Hash     []byte `json:"hash"`
	Duration int64  `json:"duration"`
}

type Profile struct {
	gorm.Model
	ProfileID int64  `json:"profile_id"`
	Name      string `json:"name"`
	Hash      []byte `json:"hash"`
	Duration  int64  `json:"duration"`
}

type ProfileViewingInfo struct {
	gorm.Model
	MediaID   int64 `json:"media_id"`
	ProfileID int64 `json:"profile_id"`
	TimePoint int64 `json:"time_point"`
	Timestamp int64 `json:"timestamp"`
}
