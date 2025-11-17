package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi dengan nol")
	}
	return nilai / pembagi, nil
}

func main() {
	// err := errors.New("this is a custom error")
	// fmt.Println(err)
	result, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("Hasil:", result)
	} else {
		fmt.Println("Terjadi kesalahan:", err.Error())
	}
	
}