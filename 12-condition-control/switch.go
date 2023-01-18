package main

import (
	"fmt"
)

func returnDayOfWeek(day int) string {
	switch day {
	case 1:
		return "Sunday (domingo)"
	case 2:
		return "Moonday (segunda)"
	case 3:
		return "Tuesday (terça-feira)"
	case 4:
		return "Wednesday (quarta-feira)"
	case 5:
		return "Thursday (quinta-feira)"
	case 6:
		return "Friday (sexta-feira)"
	case 7:
		return "Saturday (sábado)"
	default:
		return "Invalid day!"
	}
}

func studentStatus(points int) string {
	switch {
	case points < 5:
		return "student reproved"
	case points > 5:
		return "student approved"
	default:
		return "invalid points"
	}
}

func main() {
	fmt.Println("Switch")
	day := returnDayOfWeek(5)
	studentPoint := studentStatus(6)

	fmt.Println(day)
	fmt.Println(studentPoint)
}
