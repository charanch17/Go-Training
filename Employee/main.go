package main

import (
	"fmt"

	helpers "Employee.com/charan/emp/Helpers"
	utils "Employee.com/charan/emp/Utils/FlatFileOperations"
)

func main() {
	utils.Write("./GeneratedFiles/EmployeeData.txt", "txt", "")
	utils.Write("./GeneratedFiles/EmployeeData.csv", "csv", "")
	utils.Write("./GeneratedFiles/EmployeeData.xlsx", "xlsx", "sheet1")
	employee := helpers.SearchEmployeeUsingEmployeeId(1)
	if employee.EmployeeId != 0 {
		fmt.Printf("%+v\n", employee)
	} else {
		fmt.Println("no employee found")
	}
}
