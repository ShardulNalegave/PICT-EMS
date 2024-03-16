package csv

type PeopleRecord struct {
	RegistrationID string `csv:"Registration ID"`
	Name           string `csv:"Name"`
}
