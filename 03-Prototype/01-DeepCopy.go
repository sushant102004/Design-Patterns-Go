/*
	Deep Copy is the process of copying everything including values from pointers not pointers itself.
*/

package main

// type Address struct {
// 	City  string
// 	State string
// }

// type Person struct {
// 	Name    string
// 	Address *Address
// }

// func main() {
// 	sushant := Person{
// 		Name: "Sushant",
// 		Address: &Address{
// 			City:  "Panipat",
// 			State: "Haryana",
// 		},
// 	}

// 	jane := sushant

// 	// We can't directly change is here as it is pointer and will change actual
// 	// value coming from sushant rather than a copy.
// 	// jane.Address.City = "Pune"

// 	// Solution - Deep Copy

// 	jane.Address = &Address{
// 		City:  sushant.Address.City,
// 		State: sushant.Address.State,
// 	}

// 	jane.Name = "Jane"
// 	jane.Address.City = "Pune"

// 	fmt.Println(sushant, sushant.Address.City)
// 	fmt.Println(jane, jane.Address.City)
// }
