package main

import "fmt"

func main() {

	fmt.Println(greet())
	fmt.Println(hello("Mayur"))
	add(5, 10)
	quotient, remainder := divisions(20, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)

	area, parameter := ractangleArea(10, 5)
	fmt.Printf("Area: %d, Parameter: %d\n", area, parameter)

	sum := addition(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d\n", sum)

}

// Function without Parameter
func greet() string {
	return "Hello from functions.go!"
}

// Function with Parameter
func hello(name string) string {
	return "Hello " + name
}

// Function with Parameters
func add(a int, b int) {
	fmt.Println(a + b)
}

// Multiple Return Values
func divisions(p int, q int) (int, int) {
	quotient := p / q
	remainder := p % q
	return quotient, remainder
}

// Named Return Values
func ractangleArea(length, width int) (area, parameter int) {

	area = length * width
	parameter = 2 * (length + width)
	return
}

// Variadic Function
func addition(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
