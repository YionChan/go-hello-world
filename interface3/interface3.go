//Embedding interfaces
//
//Although go does not offer inheritance, it is possible to create a new interfaces by embedding other interfaces.
package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

//Embedding interfaces
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func main() {

	e := Employee{
		firstName:   "peng",
		lastName:    "chan",
		basicPay:    20000,
		pf:          500,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left = ", empOp.CalculateLeavesLeft())

}
