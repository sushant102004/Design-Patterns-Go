package main

import "fmt"

/*
	Benefits of using interface factory.
	We get an additional layer of abstraction because now we can't change underlying data directly.
*/

type PersonInterface interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello, How are you?")
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("I'm too tired to say hello")
}

func MakeNewPerson(name string, age int) PersonInterface {
	if age > 100 {
		return &tiredPerson{
			name: name,
			age:  age,
		}
	} else {
		return &person{
			name: name,
			age:  age,
		}
	}
}
