package sessions

import "time"

type Session struct {
	RegistrationID string    `json:"registration_id"`
	EntryTime      time.Time `json:"entry_time"`
}
