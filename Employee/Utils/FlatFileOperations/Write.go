package utils

import (
	"strconv"

	helpers "Employee.com/charan/emp/Helpers"
)

func getStrRows(data [][]string) string {
	finalString := ""
	for _, row := range data {
		for index, col := range row {
			finalString += col
			if index < len(row)-1 {
				finalString += ","
			}
		}
		finalString += "\n"
	}
	return finalString
}

func Write(FileName, FileType, SheetName string) {
	data := [][]string{}
	HeaderRow := []string{"EmployeeId", "Name", "Age", "Designation", "Salary"}

	for _, employee := range helpers.GetAllEmployees() {
		data = append(data, []string{strconv.Itoa(employee.EmployeeId), employee.FirstName + employee.LastName, strconv.Itoa(employee.Age), employee.Designation, strconv.Itoa(employee.Salary)})
	}

	switch FileType {
	case "txt":
		finalString := getStrRows(data)
		HeaderStr := ""
		for index, col := range HeaderRow {
			HeaderStr += col
			if index < len(HeaderRow)-1 {
				HeaderStr += ","
			}
		}
		writer := TextFileWriter{data: finalString, FileName: FileName, Header: HeaderStr + "\n"}
		writer.WriteToFile()

	case "csv":
		writer := CsvWriter{FileName: FileName, Header: HeaderRow, Data: data}
		writer.WriteToFile()
	case "xlsx":
		writer := ExcelWriter{FileName: FileName, SheetName: SheetName, HeaderRow: HeaderRow, Data: data}
		writer.WriteToFile()

	}
}
