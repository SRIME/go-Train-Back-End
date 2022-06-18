package main

import "fmt"

func main() {
	fmt.Println("-----If else-----")

	marks := 45

	if marks < 45 {
		fmt.Println("You Fail")
		return
	} else if marks < 60 {
		fmt.Println("Pass")
		return
	} else if marks < 85 {
		fmt.Println("Distinct")
		return
	}
	if marks < 100 {
		fmt.Println("Topper")
		return
	}

	//declare varibles in the if statement itself
	if num := 5; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than 10")
	}

	//Commonly used in web request and response
	// if err != nil {
	// 	return nil, err
	// }
}
