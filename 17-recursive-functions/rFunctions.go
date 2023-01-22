package main

import "fmt"

func fibonacci(value uint) uint {
	if value <= 1 {
		return value
	}

	return fibonacci(value-2) + fibonacci(value-1)
}

func main() {
	position := uint(12)

	for iterator := uint(8); iterator < position; iterator++ {
		fmt.Println(iterator)
	}

	fmt.Println(fibonacci(position))
}
