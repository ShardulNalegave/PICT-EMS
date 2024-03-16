package models

type Student struct {
	RegistrationID string `gorm:"primaryKey"`
	Name           string
}
