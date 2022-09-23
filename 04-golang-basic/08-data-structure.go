package main

import "fmt"

func main() {
	// Map
	var animals map[string]int
	animals = map[string]int{}

	animals["Lion"] = 0
	animals["Shark"] = 1

	fmt.Println("Lion:", animals["Lion"])
	fmt.Println("Shark:", animals["Shark"])

	// Array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
}
