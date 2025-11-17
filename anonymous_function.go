package main

import "fmt"

type Blacklist func(string) bool

func blacklist(name string) bool {
	return name == "Anjing"
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklisted,", name)
	} else {
		fmt.Println("Welcome,", name)
	}
}

func main(){
	registerUser("Aji", blacklist)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
	
}