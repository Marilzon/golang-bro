package main

import (
	"fmt"
	"time"
)

func main() {
	iterator := 0

	for iterator < 10 {
		iterator++
		fmt.Println("Incrementing +1 to iterator", iterator)
		time.Sleep(time.Second)
	}

	fmt.Println(iterator)
}
