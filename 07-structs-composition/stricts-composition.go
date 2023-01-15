package main

import "fmt"

type Vehicle struct {
	whells int8
}

type Car struct {
	Vehicle
	name string
}

func main() {
	println("Struct composotion")

	corsa := Car{Vehicle{4}, "Corsa"}
	fmt.Println((corsa))
}
