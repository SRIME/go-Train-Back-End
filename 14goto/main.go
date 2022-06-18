package main

import "fmt"

func main() {
	fmt.Println("-----Go to Statement-----")

	index := 4

	if index == 5 {
		goto jump
	}
	fmt.Println("No Jump")
jump: //jump sequentially executes even if it was not jumped from goto statement
	fmt.Println("Hello from Jump")
}
