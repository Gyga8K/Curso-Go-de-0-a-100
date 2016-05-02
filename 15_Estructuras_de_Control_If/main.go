package main

import (
	"fmt"
)

func main() {

	//Estructuras de Control: if

	// if 5 < 6 {
	// 	fmt.Println("5 es menor que 6")
	// }

	a := 6

	if a < 6 {
		fmt.Println("a es menor que 6")
	} else if a > 6 {
		fmt.Println("a es mayor que 6")
	} else {
		fmt.Println("a es igual 6")
	}

	fmt.Println("Fin del programa")

}
