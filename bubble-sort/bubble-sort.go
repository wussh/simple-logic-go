package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	numbers := []int{64, 25, 12, 22, 11}

	fmt.Println("Sebelum Sorting:", numbers)

	bubbleSort(numbers)

	fmt.Println("Setelah Sorting:", numbers)
}
