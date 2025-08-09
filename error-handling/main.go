package main

import "fmt"

func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil

}

func main() {
	fmt.Println("started Error handling in functions")

	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Println("division of 2 numbers is", ans)
}
