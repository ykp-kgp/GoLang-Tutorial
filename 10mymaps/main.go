package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	language := make(map[string]string)
	language["JavaScript"] = "JS"
	language["Ruby"] = "RB"
	language["Python"] = "PY"

	fmt.Println("List of all language", language)

	fmt.Println(language["Ruby"])

	delete(language, "Ruby")
	fmt.Println("List of all language", language)
	language["R"] = "R"
	// Loops in go lang

	for key, value := range language {
		fmt.Printf("For key %v the value is %v\n", key, value)
	}

}
