package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

// The main function returning a print line message
func main() {
	fmt.Println("Runnig main function")
	auxiliar.Write()

	error := checkmail.ValidateFormat("maxmaril@hotmail.com")
	fmt.Println(error)
}
