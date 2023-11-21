/*
	Subscribe & Listen
*/

package main

import (
	"fmt"
)

type Observable struct {
	Subs []Observer
}

func (o *Observable) Subscribe(x Observer) {
	o.Subs = append(o.Subs, x)
}

func (o *Observable) Unsubscribe(x Observer) {
	newList := []Observer{}

	for _, sub := range o.Subs {
		if sub == x {
			continue
		} else {
			newList = append(newList, sub)
		}
	}
	o.Subs = newList
	fmt.Println("Successfully Unsubscribed")
}

func (o *Observable) CheckSubList() {
	for _, sub := range o.Subs {
		fmt.Println(sub)
	}
}

func (o *Observable) FireEvent(data interface{}) {
	for _, sub := range o.Subs {
		sub.Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Person struct {
	Observable
	Name string
}

func NewPerson(name string) Person {
	return Person{
		Observable: Observable{},
		Name:       name,
	}
}

func (p *Person) GotHeartAttack() {
	p.FireEvent("Heart Attack Emergency: " + p.Name)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Println(data.(string))
}

func main() {
	p := NewPerson("John")
	doctorService := &DoctorService{}
	p.Subscribe(doctorService)

	p.CheckSubList()

	p.GotHeartAttack()

	p.Unsubscribe(doctorService)

	p.CheckSubList()

}
