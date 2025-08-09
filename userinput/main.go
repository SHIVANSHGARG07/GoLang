package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey tell name")

	/* scan function //
	// 1) it only takes and read value upto white space
	// 2) so we can use bufio if need to read sentence
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello, Mr:", name)
	*/

	/*
		// bufio function //
		1)uses reader and standrad input type that OS reads or use
		2)we can also tell when to stop using read string method

	*/

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)

}
