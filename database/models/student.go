package models

type Student struct {
	RegistrationID string `gorm:"primaryKey" json:"registration_id"`
	Name           string `json:"name"`
}
