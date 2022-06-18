package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //reader

	//Get input from cl
	fmt.Println("Enter the first number: ")
	firstNum, _ := reader.ReadString('\n') //input first string

	fmt.Println("Enter the second number: ")
	secondNum, _ := reader.ReadString('\n') //input second string

	//convert string to integer
	intFirst, err1 := strconv.ParseFloat(strings.TrimSpace(firstNum), 64)
	intSecond, err2 := strconv.ParseFloat(strings.TrimSpace(secondNum), 64)

	intSum := intFirst + intSecond

	if err1 != nil {
		fmt.Println(err1)
	} else if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("")
		fmt.Println("Sum = ", intSum)
	}
}

//wrong because the firstNum contains the '\n' which is a whitespace, hence whitespace cannot be converted
// intFirst, err := strconv.ParseFloat(firstNum, 64)
