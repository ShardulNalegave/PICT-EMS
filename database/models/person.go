package models

type PersonKind string

const (
	StudentPersonKind PersonKind = "STUDENT"
	StaffPersonKind   PersonKind = "STAFF"
)

type Person struct {
	RegistrationID string     `gorm:"primaryKey" json:"registration_id"`
	PersonKind     PersonKind `json:"person_kind"`
}
