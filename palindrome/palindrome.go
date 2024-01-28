package main

import "fmt"

func isPalindrome(s string) bool {
	runes := []rune(s)

	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-i-1] {
			return false
		}
	}

	return true
}

func main() {
	palindromeString := "madam"
	nonPalindromeString := "hello"

	if isPalindrome(palindromeString) {
		fmt.Println(palindromeString, "is a palindrome.")
	} else {
		fmt.Println(palindromeString, "is not a palindrome.")
	}

	if isPalindrome(nonPalindromeString) {
		fmt.Println(nonPalindromeString, "is a palindrome.")
	} else {
		fmt.Println(nonPalindromeString, "is not a palindrome.")
	}
}
