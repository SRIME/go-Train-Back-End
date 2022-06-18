package main

import "fmt"

func main() {
	fmt.Println("Struct: ")
	sri := User{"Srikarthick", "sri@gmail.com", true, 20}
	fmt.Println("Struct with only values: ", sri)
	fmt.Printf("Struct with key and value format: %+v\n", sri)
	fmt.Printf("Name: %v, e-mail: %v\n", sri.Name, sri.Email)

	sri.getEmail() //gets old email initialized by sri object
	sri.newEmail() //gets the reference of the newly initialized e-mail to change the actual e-mail present in the object use pointer
	fmt.Println("It didnt change at all(use pointers for permanent change): ", sri.Email)
}

type User struct { //uppercase to acces it publicly
	Name       string //uppercase
	Email      string
	IsLoggedIn bool
	Age        int
}

//Methods are the functions inside or related to struct or class

//trying to access a value of struct using method
func (u User) getEmail() {
	fmt.Println("Email present value of the e-mail: ", u.Email)
}

func (u User) newEmail() {
	u.Email = "pre2pro@gmail.com"
	fmt.Println("Newly Initialized e-mail: ", u.Email)
}
