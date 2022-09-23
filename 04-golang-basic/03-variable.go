package main

import "fmt"

func main() {
	// Variables Declaration
	var firstName string = "John"

	var lastName string
	lastName = "Wick"

	fmt.Printf("Halo %s %s\n", firstName, lastName)

	// Declaration Variables without type data
	firstName2 := "John"
	lastName2 := "Wick"
	fmt.Printf("Halo %s %s\n", firstName2, lastName2)

	// Declaration Multi Variables
	var first, second, third string
	first, second, third = "Satu", "Dua", "Tiga"

	seventh, eighth, ninth := "Tujuh", "Delapan", "Sembilan"

	fmt.Println(first, second, third, seventh, eighth, ninth)

	// Declaration Underscore Variables
	_ = "Belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "John", "Wick"
	fmt.Println(name)

	// Constants
	const firstName3 string = "John"
	fmt.Print("Halo ", firstName3, "!\n")
}
