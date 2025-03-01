package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Variable")

	//Data Types and their usage
	// var username string = "Yogesh"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T \n", username)

	// var isLoggedIn bool = false
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var smallInt uint8 = 255
	// fmt.Println(smallInt)
	// fmt.Printf("Variable is of type: %T \n", smallInt)

	// var smallFloat float32 = 420.12345698745632145698745
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n", smallFloat)

	// var largeFloat float64 = 420.12345698745632145698745
	// fmt.Println(largeFloat)
	// fmt.Printf("Variable is of type: %T \n", largeFloat)

	// Default Value and Aliases
	// var anotherVariable float32
	// fmt.Println(anotherVariable)
	// fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	// var userName = "username@gmail.com"
	// fmt.Println(userName)
	// fmt.Printf("The type of the variable is %T \n", userName)

	// no var style

	// numberOfUsers := 300000
	// fmt.Println(numberOfUsers)

	foo, bar := 40, "Hemlo Domst"
	foo, jazz := 60, 10 // valid because jazz is new declaration
	fmt.Println(foo)
	fmt.Println(jazz)
	fmt.Println(bar)

	// you can't use walrus operator (:=) more than twice for a single variable because it is declaration + assignment
	// You cannot use := outside func
}
