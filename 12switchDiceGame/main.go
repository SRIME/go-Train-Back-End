package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case with Dice Game")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 //Generate only between 0 to 6

	fmt.Println("Value on Dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can start playing")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough //executes next case when this case is executed
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6 and can play again")
	default:
		fmt.Println("Wrong number try again!!!")
	}
}
