package ch04

import "time"

// Employee type
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// AwardAnnualRaise function
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
