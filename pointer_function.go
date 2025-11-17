package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func ChangeCountryToIndonesia(address Address) { // pass by value
// 	address.Country = "Indonesia"
// }
func ChangeCountryToIndonesia(address *Address) { // pass by reference
	address.Country = "Indonesia"
}

func main(){
	// var address1 *Address = &Address{"Jakarta", "DKI Jakarta", ""} // cara lain membuat pointer
	address1 := Address{"Jakarta", "DKI Jakarta", ""} // pointer ke Address literal
	ChangeCountryToIndonesia(&address1)
	fmt.Println(address1)
}
