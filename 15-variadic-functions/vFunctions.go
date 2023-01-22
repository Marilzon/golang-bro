package main

import "fmt"

func sumValues(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}

func main() {
	userEntry := sumValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(userEntry)
}
