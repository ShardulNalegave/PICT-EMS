package tsdb

import "time"

type Session struct {
	RegistrationID string    `json:"registration_id"`
	Location       string    `json:"location"`
	EntryTime      time.Time `json:"entry_time"`
	ExitTime       time.Time `json:"exit_time"`
}
