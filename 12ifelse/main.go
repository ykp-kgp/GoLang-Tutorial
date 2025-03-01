package main

import "fmt"

func main() {
	fmt.Println("If else in GoLang")

	loginCont := 2

	var result string
	if loginCont < 10 {
		result = "Regular User"
	}

	fmt.Println(result)
}
