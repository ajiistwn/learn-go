package main

import "fmt"

func main() {
	
	var (
		firstName string = "Aji"
		lastName  string = "Setiawan"
	)
	var fullName string = firstName + " " + lastName

	var umur int8 = 25

	negara := "Indonesia"

	fmt.Println(negara)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)
	fmt.Println(umur)

	var lengthOfFirstName int = len(firstName)
	var firstChar byte = firstName[0]

	fmt.Println(lengthOfFirstName)
	fmt.Println(firstChar)
}