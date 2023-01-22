package main

import "fmt"

func restoreExec() {
	if r := recover(); r != nil {
		fmt.Println("Exec rescue success!")
	}
}

func createPanicCase(x, y float64) bool {
	defer restoreExec()

	average := (x + y) / 2

	if average > 7 {
		return true
	} else if average < 6 {
		return false
	}

	panic("Wrong case, panic invoked!")
}

func main() {
	fmt.Println(createPanicCase(7, 7))
	fmt.Println("Post of run!")
}
