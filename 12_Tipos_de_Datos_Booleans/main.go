package main

import (
	"fmt"
)

func main() {

	//Dato tipo bool
	//2 posibles valores
	//true y false

	var resultado bool

	resultado = 5 < 6
	fmt.Println(" 5 < 6 - ", resultado)
	fmt.Println("***************************************")
	resultado = (5 < 6) && (3 > 1)
	fmt.Println(" (5 < 6) && (3 > 1) - ", resultado)
	fmt.Println("***************************************")
	resultado = (5 > 6) || (3 > 1)
	fmt.Println(" (5 > 6) || (3 > 1) - ", resultado)
	fmt.Println("***************************************")
	fmt.Println("!resultado", !resultado)

}
