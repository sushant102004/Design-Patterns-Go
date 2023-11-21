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

// Example for Property Change Notification
type PropertyChange struct {
	Name  string
	Value interface{}
}

type Driver struct {
	Observable
	age int
}

func NewDriver(age int) Driver {
	return Driver{
		Observable: Observable{
			Subs: make([]Observer, 0),
		},
		age: age,
	}
}

func (d *Driver) Age() int { return d.age }
func (d *Driver) SetAge(age int) {
	if age == d.age {
		return
	}
	d.age = age
	d.FireEvent(PropertyChange{
		Name:  "age",
		Value: age,
	})
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	// pc -> PropertyChange{}
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "age" {
			if pc.Value.(int) >= 18 {
				t.o.Unsubscribe(t)
				fmt.Println("You can drive now.")
			} else {
				fmt.Println("Driver age updated.")
			}
		}
	}
}

func main() {
	// p := NewPerson("John")
	// doctorService := &DoctorService{}
	// p.Subscribe(doctorService)

	// p.CheckSubList()

	// p.GotHeartAttack()

	// p.Unsubscribe(doctorService)

	// p.CheckSubList()

	driver := &Driver{
		age: 16,
	}

	trafficMgmt := &TrafficManagement{
		o: driver.Observable,
	}

	driver.Subscribe(trafficMgmt)

	driver.SetAge(17)

	driver.SetAge(18)
}
