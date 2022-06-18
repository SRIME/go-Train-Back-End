package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	var ptr1 *int
	fmt.Println("Pointer: ", ptr1)
	myNum := 10

	var ptrNum = &myNum                       // Address of myNum is assigned to ptrNum
	fmt.Println("value of myNum: ", *ptrNum)  //pointer of ptrNum returns the value stored in address
	fmt.Println("Address of myNum: ", ptrNum) //ptrNum contains the address

	// change the value in the memory address
	*ptrNum = *ptrNum + 10
	fmt.Println("New value of myNum", myNum)
	fmt.Println("New value of ptrNum", *ptrNum)

}
