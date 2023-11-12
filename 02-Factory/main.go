package main

type Person struct {
	Name      string
	Age       int
	EyesCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:      name,
		Age:       age,
		EyesCount: 2,
	}
}
