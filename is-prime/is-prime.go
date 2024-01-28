package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(numbers []int) []int {
	var primes []int
	for _, num := range numbers {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func main() {
	var numbers [1000]int

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}

	sliceNumbers := numbers[:]

	primeNumbers := findPrimes(sliceNumbers)

	fmt.Println(primeNumbers)
}
