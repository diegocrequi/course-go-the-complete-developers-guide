package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jimparty@gmail.com",
			zipCode: 41009,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
