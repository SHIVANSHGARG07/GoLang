package main

import (
	"fmt"
	// "io/ioutil"
	"os"
)

func main() {

	/**
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file", err)
		return
	}

	defer file.Close()

	content := "hello bro"
	// io return int,error
	// _, error := io.WriteString(file, content+"\n")
	// if error != nil {
	// 	fmt.Println("Error while writing data in file", error)
	// 	return
	// }

	// to check how m uch bytes it return
	byte, error := io.WriteString(file, content+"\n")
	fmt.Println(byte)
	if error != nil {
		fmt.Println("Error while writing data in file", error)
		return
	}

	*/

	//////////////////2/////////////

	/*
		// from here starts buffer
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file", err)
			return
		}

		defer file.Close()

		// buffer is also called as temp storage
		// create a buffer to read the file content
		buffer := make([]byte, 1024)

		for {
			n, err := file.Read(buffer)

			// if reached end while reading file
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("Error while reading file", err)
			}

			// Process the file reading
			// read buffer string from 0 to n size
			fmt.Println(string(buffer[:n]))
		}
	*/

	///////////////3///////////
	// there is alternate way of using io util package
	// it provides Readfile function
	// but this io util is req when we need to load full file in memory and read from there
	// this doesnt work like buffer

	// content, err := ioutil.ReadFile("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while reading", err)
	// 	return
	// }
	// fmt.Println((string(content)))

	// this also works if wrote as os.ReadFile
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading", err)
		return
	}
	fmt.Println((string(content)))

}
