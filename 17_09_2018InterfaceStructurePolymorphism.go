package main

import "fmt"

type SalaryCalculator interface {
	calculateSalary() float64
}
type Permanent struct {
	empId    int
	basicPay int
	pf       int
}
type Contract struct {
	empId    int
	basicPay int
}

func (c Contract) calculateSalary() float64 {
	return float64(c.basicPay)
}

func (p Permanent) calculateSalary() float64 {
	return float64(p.basicPay + p.pf)
}
func main() {

	pemp1 := Permanent{1, 232, 4545}
	pemp2 := Permanent{1, 23545, 4545}
	cont1 := Contract{1, 4545}
	calucator := []SalaryCalculator{pemp1, pemp2, cont1}
	expense(calucator)
}
func expense(calculators []SalaryCalculator) {
	exp := 0.0
	for _, value := range calculators {
		exp = exp + value.calculateSalary()
	}
	fmt.Println("Expense = ", exp)
}
