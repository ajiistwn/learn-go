package main

import "fmt"

func main(){
	names := [...]string{"Aji", "Setiawan", "Jhon", "Diana", "Eve", "Frank"}

	slice1 := names[4:6]
	fmt.Println("Slice 1:", slice1)

	slice2 := names[:3]
	fmt.Println("Slice 2:", slice2)

	slice3 := names[3:]
	fmt.Println("Slice 3:", slice3)

	slice4 := names[:]
	fmt.Println("Slice 4:", slice4)

	var slice5 []string = names[:]// convert array to slice
	fmt.Println("Slice 5:", slice5)

	fmt.Println("Length of Slice 5:", len(slice5))
	fmt.Println("Capacity of Slice 5:", cap(slice5))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "sabtu modif"
	daysSlice1[1] = "minggu modif"
	
	fmt.Println("Days Array:", days)
	fmt.Println("Days Slice 1:", daysSlice1)

	daysSlice2 := append(daysSlice1, "libur baru")
	daysSlice2[0] = "ups"
	fmt.Println("Days Slice 2:", daysSlice2)
	fmt.Println("Days:", days)

	// make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Aji"
	newSlice[1] = "Setiawan"
	newSlice = append(newSlice, "Jhon")
	fmt.Println("New Slice:", newSlice)
	fmt.Println("Length of New Slice:", len(newSlice))
	fmt.Println("Capacity of New Slice:", cap(newSlice))

	// copy slice
	originalSlice := []string{"Aji", "Setiawan", "Jhon"}
	copiedSlice := make([]string, len(originalSlice), cap(originalSlice))
	copy(copiedSlice, originalSlice)
	fmt.Println("Original Slice:", originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	//declare and initialize slice
	iniArray := [...]string{"Aji", "Setiawan", "Jhon"} //atau [5]string{"Aji", "Setiawan", "Jhon"}
	iniSlice := []string{"Aji", "Setiawan", "Jhon"}
	fmt.Println("Array:", iniArray)
	fmt.Println("Slice:", iniSlice)
}
