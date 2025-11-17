package main

import "fmt"

func main() {
	const firstName string = "Aji"
	const lastName string = "Setiawan"
	const fullName string = firstName + " " + lastName

	const (
		umur     int8   = 25
		negara   string = "Indonesia"
	)

	// fullName = "Budi Sant" // error: cannot assign to fullName

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)
}
