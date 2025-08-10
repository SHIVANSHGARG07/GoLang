package main

import "fmt"

func modify(num *int) {
	*num = *num * 20
	*num = *num + 20
}
func main() {
	var num int
	num = 2

	// this *int represent that what type of data should this pointer point to
	var ptr *int
	ptr = &num

	fmt.Println(ptr)

	num = 4
	fmt.Println(num)
	fmt.Println(ptr)

	// default value of ptr is nil
	var pointer *int
	if pointer == nil {
		fmt.Println("pointer is nil")
	}

	// how to pass pointer in functions
	val := 10
	modify(&val)
	fmt.Println("modify value", val)

}
