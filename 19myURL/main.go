package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbf9s"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myURL)

	// parsing the url

	result, _ := url.Parse(myURL)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.devs",
		Path:    "payment",
		RawPath: "user=yogesh",
	}

	anotherURL := partofURL.String()
	fmt.Println(anotherURL)

}
