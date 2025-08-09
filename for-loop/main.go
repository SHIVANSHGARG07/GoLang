// it only has for loop
// no while loop , no do while

package main

import "fmt"

func main() {
	// basic loop
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	// infinite loop

	counter := 0
	for {
		fmt.Println("infinite loop", counter)
		counter++
	}
}
