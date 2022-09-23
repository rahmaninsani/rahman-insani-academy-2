package main

import "fmt"

// Struct
type person struct {
	name string
	age  int
}

// Embedded Struct
type student struct {
	grade int
	person
}

func main() {
	//var s1 student
	var s1 = student{}

	s1.name = "John Wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name:\t\t", s1.name)
	fmt.Println("age:\t\t", s1.age)
	fmt.Println("person.age:\t", s1.person.age)
	fmt.Println("grade:\t\t", s1.grade)

	// Combine Struct with Slice
	var students = []person{
		{
			name: "Alice",
			age:  22,
		},
		{
			name: "Bob",
			age:  25,
		},
	}

	for _, student := range students {
		fmt.Printf("%s age is %d\n", student.name, student.age)
	}
}
