// /*
// 	Deep Copy is the process of copying everything including values from pointers not pointers itself.
// */

package main

// import "fmt"

// type Address struct {
// 	City  string
// 	State string
// }

// func (a *Address) DeepCopy() *Address {
// 	return &Address{
// 		City:  a.City,
// 		State: a.State,
// 	}
// }

// type Person struct {
// 	Name    string
// 	Address *Address
// }

// func (p *Person) DeepCopy() *Person {
// 	newPerson := *p
// 	newPerson.Address = p.Address.DeepCopy()

// 	return &newPerson
// }

// func main() {
// 	sushant := Person{
// 		Name: "Sushant",
// 		Address: &Address{
// 			City:  "Panipat",
// 			State: "Haryana",
// 		},
// 	}

// 	jane := sushant.DeepCopy()

// 	jane.Address.City = "Pune"
// 	jane.Address.State = "Maharashtra"

// 	fmt.Println(sushant.Name, sushant.Address)
// 	fmt.Println(jane.Name, jane.Address)
// }
