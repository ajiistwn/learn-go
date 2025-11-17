package main

import "fmt"

func main(){
	var result = 10 + 10
	fmt.Println("Result before operation:", result)

	var a = 10
	var b = 20
	result = a + b
	fmt.Println("Result after addition:", result)

	result = b - a
	fmt.Println("Result after subtraction:", result)

	result = a * b
	fmt.Println("Result after multiplication:", result)

	result = b / a
	fmt.Println("Result after division:", result)

	result = b % a
	fmt.Println("Result after modulus:", result)

	result += 5
	fmt.Println("augmented assignment:", result)

	result++
	fmt.Println("increment:", result)
}
