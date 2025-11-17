package main

import "fmt"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{Message: "Id tidak boleh kosong"}
	}
	if id != "aji" {
		return &NotFoundError{Message: "Data tidak ditemukan"}
	}
	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// if e, ok := err.(*ValidationError); ok {
		// 	fmt.Println("Validation error:", e.Error())
		// } else if e, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("Not found error:", e.Error())
		// } else {
		// 	fmt.Println("Unknown error:", err.Error())
		// }

		switch e := err.(type) {
		case *ValidationError:
			fmt.Println("Validation error:", e.Error())
		case *NotFoundError:
			fmt.Println("Not found error:", e.Error())
		default:
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		fmt.Println("Data berhasil disimpan")
	}
}