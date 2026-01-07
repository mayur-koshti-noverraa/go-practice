package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var a int = 10
	var p *int = &a // p is a pointer to an integer, storing the address of a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (Address of a):", p)
	fmt.Println("Value at the address stored in p:", *p) // Dereferencing pointer p to get the value of a
}
