package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	var one int = 24
	fmt.Println(one)

	var ptr *int = &one
	fmt.Println(*ptr)
	*ptr = *ptr * 2
	fmt.Println(*ptr)
	fmt.Println(one)
}
