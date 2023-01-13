package main

import "fmt"

type EmployerExpenseCalculator interface {
	calculateIndividualExpense() int
}

type SoftwareEmployee struct {
	id                   int
	name                 string
	salary               int
	enchasedEarnedLeaves int
	bonus                int
}

func (emp SoftwareEmployee) calculateIndividualExpense() int {
	return emp.salary + emp.enchasedEarnedLeaves*int(emp.salary/30) + emp.bonus
}

type BuildingMaintainanceWorker struct {
	name      string
	dailyWage int
}

func (wrkr BuildingMaintainanceWorker) calculateIndividualExpense() int {
	return wrkr.dailyWage * 30
}

func caller(clr EmployerExpenseCalculator) {
	fmt.Println("expense is", clr.calculateIndividualExpense())

}

func main() {
	worker1 := BuildingMaintainanceWorker{name: "ram", dailyWage: 200}
	SoftwareEmployee1 := SoftwareEmployee{name: "charan", salary: 400000}
	caller(worker1)
	caller(SoftwareEmployee1)

}
