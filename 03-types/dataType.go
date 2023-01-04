package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var integer int = 1

	var integer8 int8 = 8
	var integer16 int8 = 16
	var integer32 int8 = 32
	var integer64 int8 = 64

	var realNumber uint = 10

	fmt.Println(reflect.TypeOf(integer))
	fmt.Println(reflect.TypeOf(integer8))
	fmt.Println(reflect.TypeOf(integer16))
	fmt.Println(reflect.TypeOf(integer32))
	fmt.Println(reflect.TypeOf(integer64))

	fmt.Println(realNumber)

	// RUNE IS ALIAS FOR A int32
	// BYTE IS ALIAS FOR A uint8

	var runeNumber rune = 0
	var byteNUmber byte = 8

	fmt.Println(reflect.TypeOf(runeNumber))
	fmt.Println(reflect.TypeOf(byteNUmber))

	var message string = "Hellow developer"
	fmt.Println(reflect.TypeOf(message))

	var errorMessage error = errors.New("Message of error!")
	fmt.Println(errorMessage)
}
