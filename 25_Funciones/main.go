package main

import (
	"fmt"
)

//Declarar Funciones 1

func imprimirNombre(nombre string) {
	fmt.Println("Fuera de Main")
	fmt.Println("El nombre es: ", nombre)
}

//Declarar Funciones 2
func suma(n1 int, n2 int) int {
	return n1 + n2
}

//Declarar Funciones 3
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

func main() {
	//Funciones

	//Llamar una Funcion
	imprimirNombre("Jose")
	fmt.Println("Dentro de main")

	fmt.Println(suma(25, 66))
	fmt.Println(resta(88, 66))
	//
	//Firma de una funcion
	fmt.Printf("Funcion suma: %T\n", suma)
	fmt.Printf("Funcion resta: %T\n", resta)

	//package math
	//func Sin(x float64) float64 // implemented in assembly language

}
