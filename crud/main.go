package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGet() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error in getting:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	// after this we need to un marshall this and store in to do
	// so use alternate method to directly store

	// alternate method

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(todo)

}

// post request
func performPost() {
	todo := Todo{
		UserId:    23,
		Title:     "Shivansh Garg",
		Completed: true,
	}

	// we have to send json format to api
	// convert from this data to json using marshalling
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string to reader for post request
	jsonReader := strings.NewReader(jsonString)

	//post request has 3 params
	// 1) url
	// 2) data: content Type
	// 3) reader of body

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	// now convert into readable format
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))

}
func basics2() {
	fmt.Println("Hello basics1")
}
func basics21() {
	fmt.Println("Hello basics21")
}
func main() {
	performGet()
	performPost()
}
