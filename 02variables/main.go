package main

import "fmt"

const PublicVariable = "Hello" //Public variables start with uppercase

func main() {
	var age uint8 = 20
	fmt.Println("Variable value: ", age)         //println for printing variable's value
	fmt.Printf("Variable is of type: %T\n", age) //printf for printing variable type

	var userName = "Sri karthick"
	fmt.Println("Variable value: ", userName)
	fmt.Printf("Variable is of type: %T\n", userName)

	isLoggedIn := true //Global variable cannot have valarus operator
	fmt.Println("Variable value: ", isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	fmt.Println(PublicVariable)
}
