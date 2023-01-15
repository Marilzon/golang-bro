package main

import "fmt"

type User struct {
	name    string
	age     uint8
	address Address
}

type Address struct {
	description string
}

func main() {
	fmt.Println("Structs")

	var user User
	user.name = "Marilzon"
	user.age = 30

	fmt.Println(user)

	addressExample := Address{"Street 1000"}

	secondUser := User{"Max", 30, addressExample}
	fmt.Println(secondUser)

	thirduser := User{name: "Sousa"}
	fmt.Println(thirduser)
}
