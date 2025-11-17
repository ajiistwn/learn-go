package main

import "fmt"

func main(){
	var arr = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)

	fmt.Println("First element:", arr[0])
	fmt.Println("Second element:", arr[1])
	fmt.Println("Third element:", arr[2])
	fmt.Println("Fourth element:", arr[3])
	fmt.Println("Fifth element:", arr[4])

	arr[2] = 100
	fmt.Println("Updated third element:", arr[2])

	fmt.Println("Array length:", len(arr))
	fmt.Println("Array elements using loop:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}

	var arr2 [3]string
	arr2[0] = "Go"
	arr2[1] = "is"
	arr2[2] = "fun"
	fmt.Println("String Array:", arr2)
	fmt.Println("length of String Array:", len(arr2))
}
