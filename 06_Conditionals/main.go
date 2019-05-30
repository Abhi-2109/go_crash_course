package main

import "fmt"

func main() {
	x := 15
	y := 10
	// If else statement
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	//Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("color is not red and blue")
	}

	// If with initailaization

	if a := 5; a > 8 {
		fmt.Println("Its Working its ", a)
	} else {
		fmt.Println("Its else ", a)
	}

}
