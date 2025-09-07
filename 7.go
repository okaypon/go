package main

import (
	"fmt"
)

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func CalculatorSalary(employees []Employee) (totalsalary float64, avgsalary float64) {

	if len(employees) == 0 {
		return 0, 0
	}
	totalsalary = 0
	for _, employ := range employees {
		totalsalary += employ.Salary
	}
	avgsalary = totalsalary / float64(len(employees))
	return totalsalary, avgsalary
}

func main() {
	employees := []Employee{
		{ID: 1, Name: "Саня", Position: "Директор", Salary: 230000},
		{ID: 2, Name: "София", Position: "Дизайнер", Salary: 180000},
		{ID: 3, Name: "Максим", Position: "Секретарь", Salary: 110000},
	}
	totalSalary, avgSalary := CalculatorSalary(employees)

	fmt.Printf("Общий фонд оплаты труда: %.2f", totalSalary)
	fmt.Printf(" Средняя зарплата: %.2f", avgSalary)
}
