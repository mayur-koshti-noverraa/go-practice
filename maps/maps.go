package main

import "fmt"

func main() {

	/***************** Map ****************/

	// Method 1 - Declaration
	fruitList1 := map[string]int{}

	// Method 2 - Using make()
	fruitList2 := make(map[string]int)

	// Adding key-value pairs
	fruitList3 := map[string]int{
		"Apple":  100,
		"Banana": 200,
		"Mango":  300,
	}

	fruitList1["Year"] = 2000
	fruitList1["Month"] = 12
	fruitList1["Day"] = 31

	fruitList2["Red"] = 100
	fruitList2["Green"] = 200
	fruitList2["Blue"] = 300

	print("======================================\n")

	// Print maps
	fmt.Println("Fruit List 1:", fruitList1)
	fmt.Println("Fruit List 2:", fruitList2)
	fmt.Println("Fruit List 3:", fruitList3)

	print("======================================\n")

	key, value := fruitList3["Mangos"] // Access value using key
	fmt.Printf("Key: %v | Value: %v\n", key, value)

	delete(fruitList2, "Red") // Delete key-value pair
	fruitList2["Yellow"] = 400
	fmt.Println("Fruit List 2 after deleting Red:", fruitList2)

	clear(fruitList3) // Clear all elements from map
	fmt.Println("Fruit List 3 after clearing:", fruitList3)

}
