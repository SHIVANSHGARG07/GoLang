package main

import "fmt"
import "strings"

func main() {
	data := "apple, orange, banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	// count function
	str := "one two three one two threer threeyu"
	count := strings.Count(str, "three")
	fmt.Println(count)

	// trim function
	str = "             helo        a   "
	fmt.Println(str)
	// doesnt consider in b/w space of words
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	// join function
	str1 := "shivansh"
	str2 := "garg"
	res := strings.Join([]string{str1, str2}, ":")
	fmt.Println(res)
}
