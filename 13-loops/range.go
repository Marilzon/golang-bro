package main

import "fmt"

func main() {
	users := [3]string{"Marilzon", "de", "Sousa"}
	fmt.Println(users)

	for index, value := range users {
		fmt.Println(index, ":", value)
	}

	for index, letter := range "MARILZON" {
		fmt.Println(index, string(letter))
	}
}
