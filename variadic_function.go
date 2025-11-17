package main

func main(){
	// result := 
	println("SumAll:", sum(1, 2, 3, 4, 5))

	numbers := []int{6, 7, 8, 9, 10}
	println("SumAll with slice:", sum(numbers...))

}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
