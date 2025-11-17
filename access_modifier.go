package main

import (
	"fmt"
	"learn-go/helper"
)
func main(){
	fmt.Println(helper.SayHello("Aji"))
	fmt.Println("Application:", helper.Aplication)
	// fmt.Println(helper.sayGoodbye("Aji")) // error: cannot refer to unexported name helper.sayGoodbye
	// fmt.Println("Version:", helper.version) // error: cannot refer to unexported name helper.version
}