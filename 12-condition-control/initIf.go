package main

import "fmt"

func main() {

	if valueTest := 100; valueTest > 50 {
		fmt.Println(valueTest, "> 50")
	} else {
		fmt.Println(valueTest, "< 50")
	}

}
