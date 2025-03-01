package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices!")
	var List = []int{1, 2, 1, 3}
	fmt.Println("The size of the list is", len(List))
	fmt.Printf("The type of the list is %T \n", List)
	List = append(List, 420, 645)
	// fmt.Println("The size of the list is", List)
	List[0] = 1000
	// List = append(List[2:])
	fmt.Println("This is sliced", List)

	highScore := make([]int, 3) // Just like calloc in C
	// make(dataType, size of allocation)

	highScore[0] = 0
	highScore[1] = 1
	highScore[2] = 2
	fmt.Println("The the highScore Slice looks like,", highScore)

	// How to remove element at particulat index i
	var index int = 2

	var courses = []string{"OR", "Recommender", "PDPP", "Stochastic Processes", "Game Theory", "Econometrics"}
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
