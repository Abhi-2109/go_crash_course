package main

import "fmt"

func main() {
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Decalre and assign

	arr := [3]int{1, 2, 3}

	fmt.Println(arr)
	fmt.Println(arr[1])

	// Slices

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) // Silcing
}
