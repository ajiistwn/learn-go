package main

import "fmt"

func loging(){
	fmt.Println("Selesai Memanggil Function")
}

func runAplication(value int) {

	defer loging()

	fmt.Println("Run Aplication")
	// result := 10 / value
	// fmt.Println("Result:", result)
}

func main(){
	runAplication(2)
}
