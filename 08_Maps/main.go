package main

import "fmt"

func main() {

	// Define map

	emails := make(map[string]string)

	emails["Abhi"] = "abhi@email.com"
	emails["Aditi"] = "aditi@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])
	fmt.Println("Length of Map ", len(emails))

	// Delete  from map

	delete(emails, "Abhi")
	fmt.Println(emails)

	var ff = map[string]int{
		"aa": 2, "bb": 3}
	fmt.Println(ff)
}
