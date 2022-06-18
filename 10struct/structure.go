package main

import "fmt"

func main() {
	fmt.Println("Struct: ")
	sri := User{"Srikarthick", "sri@gmail.com", true, 20}
	fmt.Println("Struct with only values: ", sri)
	fmt.Printf("Struct with key and value format: %+v\n", sri)
	fmt.Printf("Name: %v, e-mail: %v\n", sri.Name, sri.Email)
}

type User struct { //uppercase to acces it publicly
	Name       string //uppercase
	Email      string
	IsLoggedIn bool
	Age        int
}
