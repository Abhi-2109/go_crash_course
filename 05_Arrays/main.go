package main

import "fmt"

func description(a []int) {
	fmt.Println(a, "length : ", len(a), " Capacity : ", cap(a))
}

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

	// Slcing
	fmt.Println("Here we go")
	a := make([]int, 5, 5)
	description(a)
	b := append(a, 2)
	b[1] = 100
	description(b)

	fmt.Println("Just check here")
	description(a)

	for i := 1; i < 10; i++ {
		a = append(a, i)
		description(a)
	}
}
