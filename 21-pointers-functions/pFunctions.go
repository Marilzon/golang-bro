package main

import "fmt"

func revertNumber(number int) int {
	return number * -1
}

func revertNumberPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	newNumber := revertNumber(number)
	fmt.Println(newNumber)

	otherNumber := 40
	fmt.Println(otherNumber)
	revertNumberPointer(&otherNumber)
	fmt.Println(otherNumber)
}
