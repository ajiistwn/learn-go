package main

import "fmt"

func main(){
	type KTP string
	type Married bool
	 
	var noKTP KTP = "3172081502000001"
	var statusMarried Married = true

	fmt.Println(noKTP)
	fmt.Println(statusMarried)
}
