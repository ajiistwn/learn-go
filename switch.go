package main

import "fmt"

func main() {
	day := 5
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with multiple expressions

	switch char := 'B'; char {
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Println(string(char), "is a vowel")
	default:
		fmt.Println(string(char), "is a consonant")
	}

	length := 5
	switch {
	case length < 5:
		fmt.Println("Short")
	case length == 5:
		fmt.Println("Medium")
	default:
		fmt.Println("Long")
	}

}
