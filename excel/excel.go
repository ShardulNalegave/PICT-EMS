package excel

import (
	"fmt"
	"time"

	"github.com/ShardulNalegave/PICT-EMS/tsdb"
	"github.com/xuri/excelize/v2"
)

func CreateReportFile(records []tsdb.Session) string {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.SetCellValue("Sheet1", "A1", "Registration ID")
	f.SetCellValue("Sheet1", "B1", "Entry Time")
	f.SetCellValue("Sheet1", "C1", "Exit Time")

	for i, rec := range records {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), rec.RegistrationID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), rec.EntryTime.Local())
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), rec.ExitTime.Local())
	}

	name := fmt.Sprintf(
		"reports/Report - %d-%d-%d.xlsx",
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
	)

	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}

	return name
}
