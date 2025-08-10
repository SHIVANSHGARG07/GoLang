package main

import (
	"fmt"
	"time"
)

// it supports 24 hour format
func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("%T\n", currentTime)

	// it cant undeerstands dd-mm-yyyy
	// we have to write example format : 02-01-2006
	// and if time then: 15-04-05
	formatted := currentTime.Format("02-01-2006,15:04:05")
	fmt.Println(formatted)

}
