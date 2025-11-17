package main

import "fmt"

type Filter func(string) string

func main(){
	sayHello("Aji", filterName)
	sayHello("Anjing", filterName)
}

func sayHello(name string, nameFilter Filter) {
	fmt.Println("Hello,", nameFilter(name))
}

func filterName(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}
