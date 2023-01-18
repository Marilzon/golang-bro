package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)

	autoLengthArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(autoLengthArray)

	mySlice := []string{"Marilzon", "de", "Sousa"}
	fmt.Println(mySlice)

	fmt.Println(reflect.TypeOf(myArray))
	fmt.Println(reflect.TypeOf(mySlice))

	mySlice = append(mySlice, "FullStack")
	fmt.Println(mySlice)

	myNewSlice := mySlice[0:3]

	fmt.Println(myNewSlice)
}
