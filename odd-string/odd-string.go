package main

import "fmt"

func main() {
	myString := "Hello, World!"

	for i, char := range myString {
		if i%2 != 0 {
			fmt.Printf("Index: %d, Character: %c\n", i, char)
		}
	}
}
