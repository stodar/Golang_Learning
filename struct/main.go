package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	//mycontact := contactInfo{"rue A", 3}
	name := person{
		firstname: "saad",
		lastname:  "dardar",
		contactInfo: contactInfo{
			"rue A",
			3,
		},
	}
	name.updateName("wail")
	name.print()
}

func (pp *person) updateName(firstname string) {
	(*pp).firstname = firstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
