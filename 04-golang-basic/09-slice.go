package main

import "fmt"

func main() {
	// Slice
	var fruits = []string{"Apple", "Grape", "Banana", "Melon"}
	fmt.Println(fruits[0])
	fmt.Println("Len: ", len(fruits))
	fmt.Println("Cap: ", cap(fruits))

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var bbFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	bFruits[0] = "Pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)
}
