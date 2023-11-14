package main

import "fmt"

type Employee struct {
	Name     string
	Position string
	Salary   int
}

func NewEmployeeFactory(position string, salary int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:     name,
			Position: position,
			Salary:   salary,
		}
	}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 70000)
	sushant := developerFactory("Sushant")
	fmt.Println(sushant)
}
