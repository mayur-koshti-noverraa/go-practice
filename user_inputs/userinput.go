package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var age int

	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	//var name string

	fmt.Println("Enter your name:")
	//fmt.Scanln(&name)

	reader := bufio.NewReader(os.Stdin)
	var name, _ = reader.ReadString('\n')
	//name = name[:len(name)-1] // Remove the newline character
	name = strings.TrimSpace(name)

	var height float64

	fmt.Println("Enter your height in cm:")
	fmt.Scanln(&height)

	fmt.Printf("You entered age: %d\n", age)
	fmt.Printf("You entered name: %s\n", name)
	fmt.Printf("You entered height: %.2f cm\n", height)

}
