package main

import "fmt"

func main() {
	fmt.Println("Welcome to Fuctions in GoLang")
	greeter()
	greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes := proAdder(2, 4, 5, 3, 10)
	fmt.Println("ProResult is: ", proRes)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Yashu di Balle Balle!")
}
func greeterTwo() {
	fmt.Println("Mera Yasu Yasu")
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// https://www.hotstar.com/in/sports/cricket/tournaments/india-tour-of-australia-2024-25/1271329069/4th-test-aus-vs-ind-day-2/1540034219/m-711729/live/watch
// https://www.hotstar.com/in/sports/cricket/tournaments/india-tour-of-australia-2024-25/1271329069/4th-test-aus-vs-ind-day-2/1540034219/m-711729/live/watch
