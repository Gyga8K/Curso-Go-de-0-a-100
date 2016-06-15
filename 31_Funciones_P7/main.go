package main

import (
	"fmt"
)

//Funciones que retornan multiples valores

func multiplicar(numero int) (n1, n2, n3 int) {
	n1 = numero * 10
	n2 = numero * 20
	n3 = numero * 30
	return
}

func multiplicar2(numero int) (int, int, int) {
	n1 := numero * 10
	n2 := numero * 20
	n3 := numero * 30
	return n1, n2, n3
}

func retorno() (string, string) {

	return "Hola", "Mundo"
}
func main() {

	fmt.Println(multiplicar(25))
	c1, c2, c3 := multiplicar(65)
	fmt.Println(c1, c2, c3)
	fmt.Println(multiplicar2(96))
	_, d2, _ := multiplicar(66)
	fmt.Println(d2)
	fmt.Println(retorno())

}
