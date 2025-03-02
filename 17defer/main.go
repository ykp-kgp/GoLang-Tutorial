package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("\nI am learning GoLang")

	myDefer()
}

func myDefer() {
	fmt.Print("check\n")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}

}
