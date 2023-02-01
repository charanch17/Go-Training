package utils

import (
	"io"
	"os"
)

type TextFileWriter struct {
	FileName string
	Header   string
	data     string
}

func (w *TextFileWriter) WriteToFile() {
	if w.FileName == "" {
		w.FileName = "EmployeeData.txt"
	}
	file, err := os.Create(w.FileName)
	if err != nil {
		panic("unable to create file")
	}
	defer file.Close()
	if len(w.Header) > 0 {
		io.WriteString(file, w.Header)

	}
	_, err = io.WriteString(file, w.data)
	if err != nil {
		panic("unable to write to " + file.Name())
	}

}
