package main

import "fmt"

func main() {
	fmt.Println("Structs in  GoLang")
	//No classes, inheritance, No super or parent

	ykp := User{"Yogesh", "ykp.kgp@gmail.com", true, 23}

	fmt.Println(ykp)
	fmt.Println("Hey")
	fmt.Printf("Yogesh's details are: %+v\n", ykp)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
