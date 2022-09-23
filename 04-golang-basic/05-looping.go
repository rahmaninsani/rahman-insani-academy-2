package main

import "fmt"

func main() {
	//	for..range
	var fruits = [4]string{"Apple", "Grape", "Banana", "Melon"}

	for i, fruit := range fruits {
		fmt.Printf("Elemen %d: %s\n", i, fruit)
	}

	// for..loop
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for..loop break continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}
