package main

import "fmt"

func main(){

	// var person = make(map[string]string)
	// person["name"] = "Aji"
	// person["age"] = "20"
	// person["address"] = "Jakarta"
	// fmt.Println("Person:", person)

	person := map[string]string{
		"name":    "Aji",
		"age":     "20",
		"address": "Jakarta",
	}
	fmt.Println("Person:", person)

	fmt.Println("Name:", person["name"])
	fmt.Println("Age:", person["age"])
	fmt.Println("Address:", person["address"])

	person["age"] = "21"
	fmt.Println("Updated Age:", person["age"])
	fmt.Println("Person after age update:", person)
	fmt.Println("Person Length:", len(person))

	delete(person, "address")
	fmt.Println("Person after deleting address:", person)
	
	fmt.Println("Map length:", len(person))
}
