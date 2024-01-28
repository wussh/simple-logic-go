package main

import "fmt"

func reverseString(input string) string {
	strBytes := []byte(input)

	length := len(strBytes)

	for i := 0; i < length/2; i++ {
		strBytes[i], strBytes[length-i-1] = strBytes[length-i-1], strBytes[i]
	}

	reversedString := string(strBytes)

	return reversedString
}

func main() {
	originalString := "Hello, World!"
	reversedString := reverseString(originalString)

	fmt.Println("Original String:", originalString)
	fmt.Println("Reversed String:", reversedString)
}
