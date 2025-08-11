package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name`
	Age     int    `json:"age`
	IsAdult bool   `json:"is_adult`
}

func main() {
	person := Person{Name: "John", Age: 34, IsAdult: true}
	// fmt.Println(person)

	//convert person into JSON Encoding (Marshlling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in marshlling", err)
		return
	}
	// convert to string because slices,array of data initially
	fmt.Println("Json person data is", string(jsonData))

	// Decoding, unmarshlling
	var personData Person
	// it takes slice of bytes so need to pass that only not string type of data
	err = json.Unmarshal(jsonData, &personData)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Decoded again", personData)

}
