package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"} 
	address2 := &address1 // pointer ke address1
	address2.City = "Bandung"

	fmt.Println(address1) // address1 ikut berubah karena address2 adalah pointer
	fmt.Println(address2) // address2 adalah pointer

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // membuat alamat baru
	fmt.Println(address1) // address1 ikut berubah karena address2 adalah pointer
	fmt.Println(address2)
}