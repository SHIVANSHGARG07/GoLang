package main

import "fmt"

type Person struct {
	firstname string
	secname   string
	age       int
}

type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var prince Person
	// fmt.Println(prince)

	prince.age = 20
	prince.firstname = "prince"
	prince.secname = "garg"

	// fmt.Println(prince)

	// user2 := Person{
	// 	firstname: "Akash",
	// 	secname:   "gupta",
	// 	age:       21,
	// }
	// fmt.Println(user2)

	var e1 Employee
	e1.Person_Details = Person{
		firstname: "prince",
		secname:   "garg",
		age:       20,
	}
	e1.Person_Contact.Email = "prince@gmail.com"
	e1.Person_Contact.Phone = "9999999999"

	e1.Person_Address = Address{
		House: 12,
		Area:  "Rajendra nagar",
		State: "Delhi",
	}

	fmt.Println(e1)

}
