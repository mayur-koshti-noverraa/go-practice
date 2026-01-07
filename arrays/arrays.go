package main

import "fmt"

func main() {

	// array declaration without values
	var arr [5]int
	fmt.Println(arr)

	// array declaration with values
	marks := [5]int{90, 80, 70, 60, 50}
	fmt.Println(marks)

	// short hand for array declaration
	scores := [...]int{100, 90, 80, 70, 60}
	fmt.Println(scores)

	// Length of array
	fmt.Println("Length of marks array:", len(marks))

	// Accessing array elements
	fmt.Println("=============== For Loop - Array==============")
	for i := 0; i < len(marks); i++ {
		fmt.Printf("marks[%d] = %d\n", i, marks[i])
	}

	// range loop
	fmt.Println("=============== Range Loop - Array==============")
	for _, value := range scores {
		fmt.Printf("%d\n", value)
	}

	// Finding average of array elements
	fmt.Println("=============== Average of Array Elements ==============")
	sum := 0
	for _, value := range marks {
		sum += value
	}
	average := float64(sum) / float64(len(marks))
	fmt.Printf("Average marks: %.2f\n", average)

}
