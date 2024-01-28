package main

import "fmt"

func isVowel(char rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if v == char {
			return true
		}
	}
	return false
}

func main() {
	myString := "Hello, World!"

	for i, char := range myString {
		if isVowel(char) {
			fmt.Printf("Index: %d, Vowel: %c\n", i, char)
		}
	}
}
