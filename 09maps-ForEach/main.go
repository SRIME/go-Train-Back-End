package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in goLang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["sh"] = "Shell-script"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	languages["r"] = "Rust"

	fmt.Println("Map: ", languages)
	fmt.Println("Accessing individual map: ", languages["js"])

	//deletion
	delete(languages, "rb")
	fmt.Println("After Deletion", languages)

	// for each looping
	for key, value := range languages {
		fmt.Printf("Key: %v, value: %v\n", key, value)
	}

}
