package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GoLang")
	content := "This is needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myNotes.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./myNotes.txt")
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data in the file is \n", string(dataBytes))
}
