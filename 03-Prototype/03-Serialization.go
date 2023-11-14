/*
	In pevious we need to rewrite whole struct into DeepCopy method. Now we will use serialization concept to deep copy the data without writing all the
	struct fields again.


	Serialization Procedure:

	Take Struct -> Encode To Bytes -> Decode Bytes To New Struct
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	City  string
	State string
}

type Person struct {
	Name    string
	Address *Address
}

func (p *Person) DeepCopy() *Person {
	var newPerson Person

	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	_ = e.Encode(&p)

	d := gob.NewDecoder(&b)
	d.Decode(&newPerson)
	return &newPerson
}

func main() {
	sushant := Person{
		Name: "Sushant",
		Address: &Address{
			City:  "Panipat",
			State: "Haryana",
		},
	}

	jane := sushant.DeepCopy()

	jane.Address.City = "Pune"
	jane.Address.State = "Maharashtra"

	fmt.Println(sushant.Name, &sushant.Address)
	fmt.Println(jane.Name, &jane.Address)
}
