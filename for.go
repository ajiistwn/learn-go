package main

import "fmt"

func main(){
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	
	numbers := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Number:", numbers[i])
	}

	names := []string{"Aji", "Bob", "John"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	fruits := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}
	for fruit, color := range fruits {
		fmt.Printf("Fruit: %s, Color: %s\n", fruit, color)
	}
}
