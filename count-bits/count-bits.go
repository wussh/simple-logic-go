package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num >>= 1
	}
	return int32(count)
}

func main() {
	// Get input from user
	fmt.Print("Enter a number: ")
	var numStr string
	fmt.Scanln(&numStr)
	num, err := strconv.ParseUint(numStr, 10, 32)
	if err != nil {
		panic(err)
	}

	// Calculate result
	result := countBits(uint32(num))

	// Print the result
	fmt.Print(result)
}
