package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("-----Slices-----")

	var numSlice = []int{0, 1, 2, 3, 4}
	// numSlice := []int{1, 2, 3, 4}
	fmt.Println(numSlice)

	// Add values
	numSlice = append(numSlice, 5, 6, 7)
	fmt.Println(numSlice)

	// Remove Values
	numSlice = append(numSlice[1:]) //Remove zeroth index value and consider the remaining values starting from 1st index
	fmt.Println(numSlice)

	numSlice = append(numSlice[1:5]) //keep only values from 1st index and 4th index remove remaining values
	fmt.Println(numSlice)

	numSlice = append(numSlice[:2]) //keep values from 0th index to 2nd index
	fmt.Println(numSlice)

	for i := 0; i < 30; i++ {
		fmt.Printf("*")
	}
	fmt.Println("")
	// make keyword
	// it is just different method of creating a slice but the output is same

	topers := make([]int, 5)
	topers[0] = 100
	topers[1] = 90
	topers[2] = 70
	topers[3] = 80
	topers[4] = 95
	// topers[5] = 99  we cannot manually add value more than the assigned values
	topers = append(topers, 99, 67) //but we can add using append method
	fmt.Println("topers: ", topers)

	// sort
	sort.Ints(topers)
	fmt.Println("Sorted Topers: ", topers)
	fmt.Println("Are the topers sorted? ", sort.IntsAreSorted(topers))

	//Remove a specific element from slice

	series := make([]string, 0)
	series = append(series, "Pokemon", "AOT", "Naruto", "Black-clover", "Jujutsu-Kaisen", "One-Piece", "Haikyuu")
	fmt.Println("series: ", series)

	removeIndex := 4
	series = append(series[:removeIndex], series[removeIndex+1:]...)
	fmt.Println("Removed only Jujutsu-Kaisen: ", series)

}
