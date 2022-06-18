package main

import "fmt"

func main() {
	fmt.Println("-----Functions-----")
	fmt.Println("Result: ", adder(1, 2))
	fmt.Println("More Numbers adder: ", sciAdder(1, 2, 3, 4, 5, 6))
}

func adder(numOne int, numTwo int) int { //Add Only two values
	return (numOne + numTwo)
}

func sciAdder(values ...int) int { //treats it as slices of values
	total := 0

	for _, val := range values {
		total = total + val
	}
	return total
}
