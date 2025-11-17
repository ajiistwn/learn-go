package main

import "fmt"

func main(){
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println("Odd number less than or equal to 7:", i)
	}
}
