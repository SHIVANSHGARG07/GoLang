package main

import "fmt"

func main() {
	var name [5]string

	// it shows ramesh and rahul at their places and else empty strings
	name[2] = "ramesh"
	name[4] = "rahul"

	fmt.Println("name is", name)
	fmt.Printf("name Array is %q\n", name)

}
