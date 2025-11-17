package main

import "fmt"

func main() {
	var a = true
	var b = false

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("a && b:", a && b) // logical AND	
	fmt.Println("a || b:", a || b) // logical OR
	fmt.Println("!a:", !a)           // logical NOT
}
