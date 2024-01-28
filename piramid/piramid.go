package main

import "fmt"

func buildPyramid(height int) {
	for i := 1; i <= height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func main() {
	height := 5

	buildPyramid(height)
}
