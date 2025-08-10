// unordered collection of data key value [pairs]

package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := make(map[string]string)

	m1["user1"] = 10
	m2["user1"] = "first"
	m1["user2"] = 20
	m2["user2"] = "second"

	fmt.Println("u1", m1["user1"], "id", m2["user1"])
	fmt.Println("u2", m1["user2"], "id", m2["user2"])

	delete(m1, "user1")
	// fmt.Println("u1", m1["user1"], "id", m2["user1"])

	// this map always return two values
	// firstly the value and second one is boolean nature
	marks, exist := m1["user1"]
	if exist {
		fmt.Println(marks)
	} else {
		fmt.Println("no user found for this")
	}
}
