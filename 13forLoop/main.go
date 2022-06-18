package main

import "fmt"

func main() {
	fmt.Println("-----Forloop-----")
	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// fmt.Println(weekDays)

	// // Normal ForLoop
	// fmt.Println("Normal Forloop: ")
	// for i := 0; i < len(weekDays); i++ {
	// 	fmt.Println(weekDays[i])
	// }

	// // ForEach Loop
	// fmt.Println("For each Loop")
	// for day := range weekDays {
	// 	fmt.Println(weekDays[day])
	// }

	// // Another ForEach Loop
	// fmt.Println("Getting Index as well ads value in foreach loop")
	// for index, value := range weekDays {
	// 	fmt.Printf("Index is %v, Value is %v\n", index, value)
	// }

	// break vs continue
	//break
	for i, v := range weekDays {
		if i == 3 {
			break //skips everything from 3rd index
		}
		fmt.Println("", v)
	}

	// continue
	for i, v := range weekDays {
		if i == 3 {
			continue //skips only 3rd index
		}
		fmt.Println(v)
	}

}
