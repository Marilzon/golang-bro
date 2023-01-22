package main

import "fmt"

func clojure() func() {
	message := "this is a clojure message"

	function := func() {
		fmt.Println(message)
	}

	return function
}

func main() {
	message := "this is a message of main"
	fmt.Println(message)

	newFunc := clojure()
	newFunc()
}
