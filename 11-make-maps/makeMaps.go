package main

import "fmt"

func main() {
	points := make(map[string]int)

	points["max"] = 100
	points["rob"] = 88
	points["not"] = 93
	points["fas"] = 100

	point, ok := points["max"]
	fmt.Println(ok)
	fmt.Println(point)

	otherPoint, ok := points["foo"]

	fmt.Println(ok)
	fmt.Println(otherPoint)

	user := map[string]string{
		"name":    "oswald",
		"surName": "klevan",
	}

	fmt.Println(user)
}
