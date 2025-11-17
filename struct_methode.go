package main

import "fmt"

type Costumer struct {
	Name string
	Age  int
}

func (c Costumer) sayHello() string {
	return "Hello, my name is " + c.Name + " and I am " + fmt.Sprint(c.Age) + " years old."
}

func main() {
	costumer := Costumer{
		Name: "Aji",
		Age:  20,
	}
	fmt.Println(costumer.sayHello())
	// fmt.Println(costumer.sayHello())
}
