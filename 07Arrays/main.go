package main

import "fmt"

func main() {
	fmt.Println("-----Arrays----- ")

	// Array declaration
	var stringArray [5]string
	stringArray[0] = "one"
	stringArray[1] = "two"
	stringArray[2] = "three"
	// stringArray[3] = "four"
	stringArray[4] = "five"

	fmt.Println("Notice the space b/w 3 and 5 notifying the nil value", stringArray)
	fmt.Println("Lengthof string array: ", len(stringArray))
	var numArray [5]int
	numArray[0] = 1
	numArray[1] = 2
	numArray[2] = 3
	// numArray[3] = 4
	numArray[4] = 5

	fmt.Println("But different for integer values", numArray)

	var singleLineDeclaration = [5]int{1, 2, 3, 4, 5}
	fmt.Println("", singleLineDeclaration)
}
