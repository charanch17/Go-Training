package utils

import (
	"encoding/csv"
	"os"
)

type CsvWriter struct {
	FileName string
	Header   []string
	Data     [][]string
}

func (w *CsvWriter) WriteToFile() {
	if w.FileName == "" {
		w.FileName = "EmployeeData.csv"
	}
	file, err := os.Create(w.FileName)
	if err != nil {
		panic("unable to create file")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(w.Header)
	err = writer.WriteAll(w.Data)
	if err != nil {
		panic("unable to write to  " + file.Name())
	}

}
