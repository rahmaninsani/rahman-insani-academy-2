package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func main() {
	// Multiple Return Values
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	// Consecutive named function parameters
	values := add(10, 20)
	fmt.Println(values)

}
