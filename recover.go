package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error terjadi:", message)
	}
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Terjadi kesalahan")
	}
	// message := recover()
	// if message != nil {
	// 	fmt.Println("Error terjadi:", message)
	// }
}

func main(){
	runApp(true)
}
