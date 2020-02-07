//Implementing multiple interfaces
//
//A type can implement more than one interface. Lets see how this is done in the following program.
package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
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
	var s SalaryCalculator = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left = ", l.CalculateLeavesLeft())

}
