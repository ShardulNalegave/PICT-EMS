package csv

import (
	"os"

	"github.com/gocarina/gocsv"
)

const (
	STAFF_DATA    = "data/data_staff.csv"
	STUDENTS_DATA = "data/data_stud.csv"
)

func ReadRecords[T any](fp string) ([]*T, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records := make([]*T, 0)
	if err := gocsv.UnmarshalFile(f, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func ReadRecordsFromString[T any](data string) ([]*T, error) {
	records := make([]*T, 0)
	if err := gocsv.UnmarshalString(data, &records); err != nil {
		return nil, err
	}

	return records, nil
}
