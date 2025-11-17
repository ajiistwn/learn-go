package main

import "fmt"

type hasName interface {
	getName() string
}

func sayHello(v hasName) {
	fmt.Println("Hello", v.getName())
}

type person struct {
	name string
}

func (p person) getName() string {
	return p.name
}


func main(){
	person := person{name: "Aji"}
	sayHello(person)
}
