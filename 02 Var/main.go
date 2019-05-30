package main

import "fmt"

func main() {
	// Data Types
	//string
	// bool
	//int

	// Using var
	//var name  = "Abhi"
	const isCool = true

	var age int32 = 246

	var size float32 = 2.3

	// Shorthand

	name := "Brad"
	email := "brad@gmail.com"

	// You can also assign thriugh this shortcut

	name1, email1 := "Abhi", "abhi@gmail.com"

	fmt.Println(name, age, isCool, size, email, name1, email1)

	fmt.Printf("%T\n", size)

}
