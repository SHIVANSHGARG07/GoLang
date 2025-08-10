package main

import (
	"fmt"
	"net/url"
)

func main() {

	// url.parse is used to parse a string into url object
	// Scheme: indicates if http or https
	// Host: Specify the domain name and optionally the port number
	// Path,Rawquery

	myurl := "https://localhost:4200/hello/world"

	parsedUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("cant parse url", err)
		return
	}
	fmt.Printf("%T\n", parsedUrl)

	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=iamprince"

	newUrl := parsedUrl.String()

	fmt.Println(newUrl)

}
