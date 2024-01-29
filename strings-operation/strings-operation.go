package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {
	// Trim all spaces at the start and end of the string
	str = strings.TrimSpace(str)

	// Create a slice to store the modified string
	modified := make([]rune, 0, len(str))

	// Iterate through the input string
	for i := 0; i < len(str); i++ {
		// Check if the character is a digit
		if str[i] >= '0' && str[i] <= '9' {
			continue
		}
		// Add the character to the modified slice
		modified = append(modified, rune(str[i]))
	}

	// Reverse the modified slice
	for i, j := 0, len(modified)-1; i < j; i, j = i+1, j-1 {
		modified[i], modified[j] = modified[j], modified[i]
	}

	// Convert the modified slice back to a string
	str = string(modified)

	return str
}

func main() {
	// Input
	fmt.Print("Enter a string: ")
	var str string
	fmt.Scanln(&str)

	// Calculate result
	result := ModifyString(str)

	// Output
	fmt.Print(result)
}
