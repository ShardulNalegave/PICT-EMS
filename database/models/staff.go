package models

type StaffMember struct {
	RegistrationID string `gorm:"primaryKey"`
	Name           string
}
