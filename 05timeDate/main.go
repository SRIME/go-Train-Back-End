package main

import (
	"fmt"
	"time"
)

func main() {

	presentTimeDate := time.Now()

	fmt.Println("Default syntax: ", presentTimeDate)
	fmt.Println("Month-Date-Year: ", presentTimeDate.Format("01-02-2006"))
	fmt.Println("Day: ", presentTimeDate.Format("Monday"))
	fmt.Println("Hours:Minutes:Seconds: ", presentTimeDate.Format("15:04:05"))

	customTime := time.Date(2002, time.June, 18, 9, 35, 0, 0, time.Local)

	fmt.Println("My birthday: ", customTime)
}
