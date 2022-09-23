package main

import "fmt"

func main() {
	// Numeric Decimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//	Boolean
	var exist bool = true
	fmt.Printf("exist? %t\n", exist)

	//	String
	var message = `Halo nama saya "Rahman Insani".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Println(message)
}
