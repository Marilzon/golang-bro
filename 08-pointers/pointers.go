package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int8 = 13
	var p *int8 = &x

	fmt.Println(x)
	fmt.Println(p)

	x = 10

	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))
}
