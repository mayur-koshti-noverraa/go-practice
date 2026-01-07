package main

import (
	"fmt"
	"strconv"
)

func main() {

	/******************* Integer Conversion *******************/

	fmt.Println("=============== INT CONVERSION =============")
	var firstInt int = 10

	// Original int value
	fmt.Printf("Value: %v \t | \t Type: %T \n", firstInt, firstInt)

	// Conversion from int to float32
	var firstFloat float32 = float32(firstInt)

	fmt.Printf("Value: %v \t | \t Type: %T \n", firstFloat, firstFloat)

	// Conversion from int to string
	var firstString string = strconv.Itoa(firstInt)

	fmt.Printf("Value: %v \t | \t Type: %T \n", firstString, firstString)

	/******************* Float Conversion *******************/
	fmt.Println("=============== FLOAT CONVERSION =============")

	var secondFloat float64 = 25.16

	// Original float64 value
	fmt.Printf("Value: %v \t | \t Type: %T \n", secondFloat, secondFloat)

	// Conversion from float64 to int
	var secondInt int = int(secondFloat)

	fmt.Printf("Value: %v \t | \t Type: %T \n", secondInt, secondInt)

	// Conversion from float64 to string
	var secondString string = strconv.FormatFloat(secondFloat, 'f', 1, 64)

	fmt.Printf("Value: %v \t | \t Type: %T \n", secondString, secondString) // round off => 25.16 ->  25.2, 25.15 -> 25.1

	/******************* String Conversion *******************/

	fmt.Println("=============== STRING CONVERSION =============")

	var thirdString string = "123.36"

	// Original string value
	fmt.Printf("Value: %v \t | \t Type: %T \n", thirdString, thirdString)

	// Conversion from string to int

	var thirdInt, err = strconv.Atoi(thirdString)

	fmt.Printf("Value: %v \t | \t Type: %T | \t Error: %v \n", thirdInt, thirdInt, err)

	// Conversion from string to float64
	var thirdFloat, err2 = strconv.ParseFloat(thirdString, 64)

	fmt.Printf("Value: %v \t | \t Type: %T| \t Error: %v  \n", thirdFloat, thirdFloat, err2)

	// Conversion from string to bool
	var fourthString string = "true"

	var thirdBool, err3 = strconv.ParseBool(fourthString)

	fmt.Printf("Value: %v \t | \t Type: %T| \t Error: %v  \n", thirdBool, thirdBool, err3)

	/******************* Boolean Conversion *******************/
	fmt.Println("=============== BOOLEAN CONVERSION =============")

	var firstBool bool = true

	// Original bool value
	fmt.Printf("Value: %v \t | \t Type: %T \n", firstBool, firstBool)

	// Conversion from bool to string
	var boolString string = strconv.FormatBool(firstBool)

	fmt.Printf("Value: %v \t | \t Type: %T \n", boolString, boolString)

}
