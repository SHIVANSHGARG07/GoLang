package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("start")
	// using this defer jabh yeah function khtm hoga main usse just phle execute hoga
	// in case of more than 1 differ
	//1) it works as lifo state
	//2) joh last mein gya voh phle ayega
	defer fmt.Println("middle")
	defer fmt.Println("middle2")
	fmt.Println("end")

}
