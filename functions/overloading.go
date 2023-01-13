package main

import "fmt"

func ExpenseTracker(vals ...interface{}) {
	var ExpenseList []string
	var TotalExpenses int

	for _, val := range vals {
		switch val.(type) {
		case int:
			TotalExpenses += val.(int)
		case int64:
			TotalExpenses += int(val.(int64))
		case int32:
			TotalExpenses += int(val.(int32))
		case float32:
			TotalExpenses += int(val.(float32))
		case float64:
			TotalExpenses += int(val.(float64))
		case string:
			ExpenseList = append(ExpenseList, val.(string))
		}
	}
	fmt.Println("list of expenses are ", ExpenseList)
	fmt.Println("Total expenses = ", TotalExpenses)
}

func main() {
	var expense1 int = 10
	var expense2 int16 = 20
	var expense3 int32 = 40
	var expense4 int64 = 30
	var expense5 float32 = 1.0
	var expense6 float64 = 20.2
	var expenseName1 = "choclates"
	var expenseName2 = "iceCreams"
	var expenseName3 = "softdrinks"
	var expenseName4 = "petrol"
	var expenseName5 = "shopping"
	var expenseName6 = "water"
	ExpenseTracker()
	ExpenseTracker(expense1, expense2, expense3, expenseName2, expenseName4, expenseName5)
	ExpenseTracker(expenseName1, expenseName5, expense1, expense2, expense3, expense4)
	ExpenseTracker(expense5, expense6, expenseName3, expenseName6)

}
