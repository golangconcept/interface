package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}
type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contact struct {
	empId    int
	basicPay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}
func (c Contact) CalculateSalary() int {
	return c.basicPay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {
	expense := 0

	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}

	fmt.Println("Total Expense per Month $%d", expense)
}

func main1() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	cemp1 := Contact{
		empId:    3,
		basicPay: 3000,
	}

	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 100,
		totalHours:  100,
	}
	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees)
}
