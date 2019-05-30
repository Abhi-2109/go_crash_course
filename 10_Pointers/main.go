package main

import "fmt"

func main() {
	a := 5
	var b *int
	b = &a

	fmt.Println(a, b)
	fmt.Printf("Type of the b  : %T\n", b)

	// Use * to read val from address

	fmt.Println(b, *b)

	// Change the value with Pointer

	*b = 10

	fmt.Println(a, *b)
}
