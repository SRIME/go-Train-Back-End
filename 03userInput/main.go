package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello, Welcome to goLang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //
	fmt.Println("Please Enter your name: ")

	//comma ok || comma err Syntax
	userName, _ := reader.ReadString('\n') // read string until new line is encountered by using enter button

	fmt.Println("Hello, ", userName)
}

// reader := bufio.NewReader(os.Stdin)
// input, err := reader.ReadString('\n')
// input, _ := reader.ReadString('\n')
// _ , err := reader.ReadString('\n')
