package utils

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelWriter struct {
	FileName  string
	SheetName string
	HeaderRow []string
	Data      [][]string
}

func (w *ExcelWriter) WriteToFile() {
	file := excelize.NewFile()

	_, err := file.NewSheet(w.SheetName)
	if err != nil {
		panic("unable to create sheet")
	}
	file.SetSheetRow(w.SheetName, "A1", w.HeaderRow)
	for index, row := range w.Data {
		pos := "A" + strconv.Itoa(index+2)
		file.SetSheetRow(w.SheetName, pos, &row)
		if err != nil {
			panic("unable to add row")
		}
	}
	err = file.SaveAs(w.FileName)
	if err != nil {
		panic("unable to save the file")
	}

}
