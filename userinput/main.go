package main

import "fmt"

func main() {
	fmt.Println("Hey tell name")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello, Mr:", name)
}
