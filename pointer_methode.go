package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main(){
	var man1 Man = Man{"John"}
	man1.Married()
	fmt.Println(man1.Name)

}
