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
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "abc@email.com",
	// 		pinCode: 000000,
	// 	},
	// }
	// fmt.Println(jim)         //{Jim Party {abc@email.com 0}}
	// fmt.Printf("%+v\n", jim) //{firstName:Jim lastName:Party contactInfo:{email:abc@email.com pinCode:0}}
	// jim.updateName("Jimmy")
	// jim.print() //{firstName:Jim lastName:Party contactInfo:{email:abc@email.com pinCode:0}}

	//Declaring struct with an embedded struct
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "abc@email.com",
	// 		pinCode: 000000,
	// 	},
	// }
	// //jimPointer is a pointer of type person holding the memory address were jim struct is located
	// //&jim returns a pointer/memory address where the real jim struct copy is stored and assigns it to a pointer
	// jimPointer := &jim //var jimPointer *person = &jim
	// jimPointer.updateName("Jimmy")
	// jim.print() //{firstName:Jimmy lastName:Party contactInfo:{email:abc@email.com pinCode:0}}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "abc@email.com",
			pinCode: 000000,
		},
	}
	jim.updateName("Jimmy") //This works
	jim.print()             //{firstName:Jimmy lastName:Party contactInfo:{email:abc@email.com pinCode:0}}
}

//This receiver can be called with any pointer to person. Not a person itself.
func (pointerToPerson *person) updateName(newFirstName string) {
	//pointerToPerson is holding the memory address where the person struct is stored and is of type person pointer
	//pointerToPerson refers to the original copy of the person struct in memory through which this receiver function is called
	//pointerToPerson is the pointer and *pointerToPerson is the value that the pointer is pointing to
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
