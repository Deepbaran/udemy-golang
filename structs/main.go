package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

// Defining custom struct type person
type person struct {
	//properties of the struct
	firstName string
	lastName  string
	//Embedding a strucu inside another struct
	// contact contactInfo
	contactInfo //If we name the property same as it's type, we do not need to specify it's type as the type will get inferred
}

func main() {
	//Declaring struct

	//1st way
	// alex := person{"Alex", "Anderson"} //Go assigns values in the order of definition
	// fmt.Println(alex)                  //{Alex Anderson}

	//2nd way
	// alex := person{firstName: "Alex", lastName: "Anderson"} //assigning values directly to the properties without depending on the order of definition
	// fmt.Println(alex)                                       //{Alex Anderson}

	//3rd way [Declaring a struct and then updating values individually]
	// var alex person
	// fmt.Printf("%+v\n", alex) //{firstName: lastName:}
	// fmt.Println(alex)         //{ } <- Two empty strings (filled automatically with default/zero value)
	// alex.firstName = "Alex"
	// fmt.Println(alex) //{Alex } <- Alex with one empty string for lastName
	// alex.lastName = "Anderson"
	// fmt.Println(alex) //{Alex Anderson}

	//Declaring struct with an embedded struct
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "abc@email.com",
	// 		pinCode: 000000,
	// 	},
	// }
	// fmt.Println(jim)         //{Jim Party {abc@email.com 0}}
	// fmt.Printf("%+v\n", jim) //{firstName:Jim lastName:Party contact:{email:abc@email.com pinCode:0}}

	//Declaring struct with an embedded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "abc@email.com",
			pinCode: 000000,
		},
	}
	fmt.Println(jim)         //{Jim Party {abc@email.com 0}}
	fmt.Printf("%+v\n", jim) //{firstName:Jim lastName:Party contactInfo:{email:abc@email.com pinCode:0}}
	jim.updateName("Jimmy")
	jim.print() //{firstName:Jim lastName:Party contactInfo:{email:abc@email.com pinCode:0}}
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
