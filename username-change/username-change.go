package main

import (
	"fmt"
)

/*
 * Complete the 'possibleChanges' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY usernames as parameter.
 */

func possibleChanges(usernames []string) []string {
	var result []string

	for _, username := range usernames {
		if canSwapToSmaller(username) {
			result = append(result, "YES")
		} else {
			result = append(result, "NO")
		}
	}

	return result
}

func canSwapToSmaller(username string) bool {
	// Convert the username to a slice of characters for easy manipulation
	usernameChars := []rune(username)

	// Iterate through the characters to find the first pair that can be swapped
	for i := 0; i < len(usernameChars)-1; i++ {
		if usernameChars[i] > usernameChars[i+1] {
			return true
		}
	}

	return false
}

func main() {
	// Input
	fmt.Print("Enter the number of usernames: ")
	var usernamesCount int
	fmt.Scanln(&usernamesCount)

	var usernames []string

	for i := 0; i < usernamesCount; i++ {
		fmt.Print("Enter username: ")
		var usernamesItem string
		fmt.Scanln(&usernamesItem)
		usernames = append(usernames, usernamesItem)
	}

	// Calculate result
	result := possibleChanges(usernames)

	// Output
	for i, resultItem := range result {
		fmt.Print(resultItem)

		if i != len(result)-1 {
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")
}
