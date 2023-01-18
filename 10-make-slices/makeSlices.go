package main

import "fmt"

func main() {
	mySlice := make([]float32, 10, 15)
	fmt.Println(mySlice)

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
