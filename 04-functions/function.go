package main

import "fmt"

func sum(x, y int8) (int8, int8) {
	plus := x + y
	sub := x - y

	return plus, sub
}

func main() {
	fmt.Println(sum(1, 2))
}
