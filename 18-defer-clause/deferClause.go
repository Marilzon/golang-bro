package main

import "fmt"

func funcWithDefer() {
	fmt.Println("func with defer")
}

func funcTwoWithDefer() {
	fmt.Println("func Two whithout defer")
}

func studentAverage(x, y float32) bool {
	defer fmt.Println("defer done!")

	average := (x + y) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcWithDefer()
	funcTwoWithDefer()

	studentAverage(10, 6)
}
