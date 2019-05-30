package main

import "fmt"

func main() {
	ids := []int{33, 76, 23, 56, 23, 134}

	// Loop through ids

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// If you don't want to use i

	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}
	var a = map[int]string{1: "abhi", 2: "raju"}

	for i, id := range a {
		fmt.Printf("%d %s\n", i, id)
	}

}
