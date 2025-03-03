package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Server request")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	// fmt.Println(string(content))

	var responseString strings.Builder

	fmt.Println(responseString)
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString)
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myURL = "http://localhost:8000/post"

	// Fake JSON payload

	requestBody := strings.NewReader(`
		{
			"courseName":"Let's go with goLang",
			"price":0,
			"platform":"Youtube.com"
		}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myURL = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}

	data.Add("firstName", "Yogesh")
	data.Add("middleName", "Kumar")
	data.Add("lastName", "Pandey")
	data.Add("email", "ykp@google.com")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
