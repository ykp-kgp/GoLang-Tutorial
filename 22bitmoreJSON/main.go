package main

import (
	"encoding/json"
	"fmt"
)

// Goal: Convert the data into JSON
type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERNBootcamp", 199, "LearnCodeOnline.in", "def123", []string{"fullstack", "js"}},
		{"AngularJS Bootcamp", 599, "LearnCodeOnline.in", "ghi123", nil},
	}

	// package this data as JSON data

	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
			"coursename": "ReactJS Bootcamp",
			"courseprice": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases wher you just want to add data to key value pair

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and the value is %v and Type is: %T\n", k, v, v)
	}
}
