package main

import "fmt"

func main() {
	fmt.Println("Array welcomes you")

	var fruitList [4]string
	// fruitList = {"Apple", "Orange", "Banana", "Kiwi"}
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Banana"
	fruitList[3] = "Kiwi"
	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length of Fruit List is: ", len(fruitList))

}
