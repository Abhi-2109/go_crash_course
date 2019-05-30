package main

import "fmt"
import "strconv"

// Define Person Structure

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting Method This is a value reciever

func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + ". My age is " + strconv.Itoa(p.age)
}

// hasBirthday method Pointer reciever

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init Person using Struct
	person1 := Person{firstName: "John", lastName: "Cena", city: "Bhopal", gender: "M", age: 20}
	person2 := Person{"Nikki", "Bella", "Bhopal", "F", 18}

	fmt.Println(person1)
	fmt.Println(person2)

	person2.age++
	fmt.Printf("%s is of age %d\n", person2.firstName, person2.age)

	fmt.Println(person1.greet())

	fmt.Println(person2.greet())

	person2.hasBirthday()
	fmt.Println(person2.age)

	person2.getMarried(person1.lastName)

	fmt.Println(person2.greet())
}
