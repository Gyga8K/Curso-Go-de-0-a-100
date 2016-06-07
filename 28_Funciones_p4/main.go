package main

import (
	"fmt"
)

//Variadic Functions

func sumar(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

func main() {

	fmt.Println(sumar(1, 2, 3, 4))
	fmt.Println(sumar(10, 65, 48, 75))
	fmt.Println(sumar(22, 25))
	fmt.Println(sumar(25))

}
