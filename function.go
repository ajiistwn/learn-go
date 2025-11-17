package main

import "fmt"

func main() {
	greet("Aji")
	sum := add(5, 10)
	fmt.Println("Sum:", sum)
	firstName, lastName := getFullName("Aji", "Setiawan")
	fmt.Println("Full Name:", firstName, lastName)

	newFirstName, _ := getFullName("Budi", "Santoso")
	fmt.Println("First Name:", newFirstName)

	namedFirstName, _ := getName("Charlie", "Brown")
	fmt.Println("Named First Name:", namedFirstName)
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func getName(firstName string, lastName string) (sendFirstName, sendLastName string) {
	sendFirstName = firstName
	sendLastName = lastName
	return	sendFirstName, sendLastName
}