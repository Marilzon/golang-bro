package main

import "fmt"

type User struct {
	name       string
	age        uint8
	resistance uint
}

func (u User) walk() {
	fmt.Printf("%s, %d is Walking...\n", u.name, u.age)
}

func (u *User) getResistence() {
	u.resistance = u.resistance + 1
}

func main() {
	oswald := User{"Oswald", 30, 10}
	david := User{"David", 21, 10}

	fmt.Println(oswald)
	fmt.Println(david)

	oswald.walk()
	david.walk()

	david.getResistence()
	david.getResistence()
	david.getResistence()

	fmt.Println(david)
}
