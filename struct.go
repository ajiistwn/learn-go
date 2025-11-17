package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	aji := Person{FirstName: "Aji", LastName: "Setiawan", Age: 20}
	fmt.Println("Person:", aji)
	fmt.Println("First Name:", aji.FirstName)
	fmt.Println("Last Name:", aji.LastName)
	fmt.Println("Age:", aji.Age)

	rizal := Person{
		FirstName: "Rizal",
		LastName:  "Setiawan",
		Age:       12,
	}
	fmt.Println("Person:", rizal)

}
