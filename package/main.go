package main

import (
	"fmt"
	"go_learning/package/ractangle"
)

func main() {
	fmt.Println("Hello, World!")

	ractangleArea := ractangle.Area(10, 5)
	ractanglePerimeter := ractangle.Perimeter(10, 5)

	fmt.Println("Ractangle Area:", ractangleArea)
	fmt.Println("Ractangle Perimeter:", ractanglePerimeter)

}
