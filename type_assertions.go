package main

import "fmt"

func random() interface{} {
	return "oke"
}

func main(){
	result := random()

	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInteger := result.(int) // panic
	// fmt.Println(resultInteger)

	switch result := result.(type) {
	case string:
		fmt.Println("String:", result)
	case int:
		fmt.Println("Integer:", result)
	default:
		fmt.Println("Unknown type")
	}
}