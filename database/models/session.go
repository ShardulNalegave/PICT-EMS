package models

import "time"

type Session struct {
	SessionID      string `gorm:"primaryKey" json:"session_id"`
	RegistrationID string `json:"registration_id"`
	EntryTime      time.Time
	ExitTime       time.Time
}
