package main

import "fmt"

func simpleFunc() {
	fmt.Println("Hello")
}

func add(a, b int) int {
	return a + b
}
func main() {
	simpleFunc()

	ans := add(3, 4)
	fmt.Println("add of two number", ans)
}
