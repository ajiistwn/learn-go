package main

import "fmt"

func main() {
	i := 10
	if i%2 == 0 {
		fmt.Println(i, "is an even number")
	} else {
		fmt.Println(i, "is an odd number")
	}

	j := 15
	if j%2 == 0 {
		fmt.Println(j, "is an even number")
	} else {
		fmt.Println(j, "is an odd number")
	}

	k := 7
	if k < 0 {
		fmt.Println(k, "is a negative number")
	} else if k == 0 {
		fmt.Println(k, "is zero")
	} else {
		fmt.Println(k, "is a positive number")
	}

	if m := 10; m > 10 {
		fmt.Println(m, "is greater than 10")
	} else {
		fmt.Println(m, "is not greater than 10")
	}
}
