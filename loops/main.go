package main

import "fmt"

func main() {
	// Arrays in Go
	//animals := [2]string{
	//	"dog",
	//	"cat",
	//}

	//animals[0] = "dog"
	//animals[1] = "cat"

	// Slices in Go
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "mouse")

	// For loop
	//for i := 0; i < len(animals); i++ {
	//	fmt.Printf("name of this animal is %s\n", animals[i])
	//}

	// Using ranges to iterate over slices
	for index, value := range animals {
		fmt.Printf("Animals[%d] = %s\n", index, value)
	}

	// Element removal from slice
	//animals = slices.Delete(animals, 0, 1)

	//fmt.Println(animals)
}
