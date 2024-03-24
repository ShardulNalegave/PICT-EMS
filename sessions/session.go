package sessions

import "time"

type Session struct {
	SessionID      string    `json:"session_id"`
	RegistrationID string    `json:"registration_id"`
	EntryTime      time.Time `json:"entry_time"`
	Location       string    `json:"location"`
}
