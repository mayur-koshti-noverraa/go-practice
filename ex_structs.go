package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func fnStruct() {

	user := User{
		Name: "Mayur",
		Age:  25,
	}

	fmt.Println(user)
}
