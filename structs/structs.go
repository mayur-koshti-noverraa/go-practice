package main

import "fmt"

type Student struct {
	name   string
	rollno int
	class  string
	marks  float64
}

func (s Student) getName() string {

	return s.name
}

func (s *Student) setName(newname string) {
	s.name = newname
}

func main() {

	var Student1 = Student{"Akshay", 101, "10th", 89.5}
	var Student2 Student = Student{name: "Ravi", rollno: 102, class: "10th", marks: 92.0}
	var Student3 = Student{rollno: 103, name: "Sonia", marks: 95.5, class: "10th"}

	fmt.Println("Student 1:", Student1)
	fmt.Println("Student 2:", Student2.name)
	fmt.Println("Student 3:", Student3)

	var Harshit Student = Student{name: "Harshit", rollno: 104, class: "11th"}

	fmt.Println("Student Harshit before setting name:", Harshit.getName())
	Harshit.setName("Harshit Kumar")

	fmt.Println("Student Harshit:", Harshit)
}
