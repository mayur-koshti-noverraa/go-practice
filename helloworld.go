package main

import (
	"fmt"
	"go_learning/utils"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("Hello World!")
	c := cases.Title(language.English)
	fmt.Println(c.String("Hello Go Module"))
	utils.PrintMessage("This is a message from the helper function.")
	fnStruct()

}
