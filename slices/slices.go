package main

import "fmt"

func main() {

	// Create method 1
	var marks []int

	// Create method 2: short hand
	scores := []int{90, 80, 70, 60, 50}

	fmt.Println("Marks:", marks)
	fmt.Println("Scores:", scores)

	// Create method 3: using make()
	numbers := make([]int, 5, 10) // length 5, capacity 5

	print("\n======================================\n")

	fmt.Println("Length of numbers slice:", len(numbers))
	fmt.Println("Capacity of numbers slice:", cap(numbers))
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Numbers:", numbers)

	print("\n======================================\n")

	// Append elements to slice
	numbers = append(numbers, 60, 70)
	fmt.Println("Numbers after appending:", numbers)
	fmt.Println("Length of numbers slice after appending:", len(numbers))
	fmt.Println("Capacity of numbers slice after appending:", cap(numbers))
}
