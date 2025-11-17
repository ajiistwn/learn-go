package main

import "fmt"	

func main(){
	greeter := greet 

	fmt.Println(greeter)
	fmt.Println(greeter("Aji"))

}

func greet(name string) string {
	return "Hello, " + name
}
