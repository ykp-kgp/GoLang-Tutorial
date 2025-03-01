package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to GoLang Programming"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")
	// comma ok || err ok
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is: %T \n", input)
	fmt.Println("This is error: ", err)
	fmt.Printf("Type of this error is: %T \n ", err)
}
