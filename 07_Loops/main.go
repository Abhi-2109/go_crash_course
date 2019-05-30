package main

import "fmt"

func main() {
	i := 1
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	for j := 1; j <= 10; j++ {
		fmt.Print(j, " ")
	}
	fmt.Println()
	m := 1
	for m < 100 {
		fmt.Print(m, " ")
		m++
	}
	fmt.Println()
}
