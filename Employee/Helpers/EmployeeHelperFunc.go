package helpers

import (
	"fmt"

	data "Employee.com/charan/emp/DB"
	models "Employee.com/charan/emp/Models"
)

func Catch() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func SearchEmployeeUsingEmployeeId(EmployeeId int) models.Employee {
	defer Catch()
	for _, Employee := range data.EmployeeTable {
		if EmployeeId == Employee.EmployeeId {
			return Employee
		}

	}
	return models.Employee{}

}

func GetAllEmployees() []models.Employee {
	return data.EmployeeTable
}
