package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting get res", res)
		return
	}
	// resource ko bnd krdo
	defer res.Body.Close()
	// gives http response
	fmt.Printf("%T\n", res)
	// fmt.Println("res", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading res", res)
		return
	}
	fmt.Println("res", string(data))

}
