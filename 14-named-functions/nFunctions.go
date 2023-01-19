package main

import "fmt"

func mathCalc(x, y int) (sumValues int, subValues int) {
	sumValues = x + y
	subValues = x - y

	return
}

func main() {
	sum, sub := mathCalc(10, 20)
	fmt.Println(sum, sub)
}
