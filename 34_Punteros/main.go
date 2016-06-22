package main

import (
	"fmt"
)

func incrementar(numero *int) {
	*numero++
	fmt.Println("Valor de Numero dentro de la funcion incrementar: ", *numero)
}

func main() {

	// // a es de tipo int
	// a := 25
	// fmt.Println("Valor de a: ", a)
	// fmt.Println("Dirección de a: ", &a)
	// //b es un puentero de a, por lo que es de tipo *int
	// b := &a
	// //Se imprime el contenidos de b
	// fmt.Println(b)
	// //Se imprime el contenido de de la variable a, que es donde apunta b
	// fmt.Println(*b)
	//
	// // b es del tipo *int no int
	// //b = 25 //error
	//
	// //Le asignamos un nuevo valor a (a) a traves de b
	// *b = 32
	// fmt.Println("valor de a: ", a)
	// //Incrementamos a
	// a++
	// //Imprimimos *b
	// fmt.Println("Valor al que apunta b: ", *b)
	//
	// //Valor 0 de los punteros es nil
	//
	// if b != nil {
	// 	fmt.Println("b es diferente de nil")
	// }
	//
	// //los punteros son comparables
	//
	// c := &a
	//
	// if b == c {
	// 	fmt.Println("b y c son iguales")
	// }
	//
	// // //Utilizar new()
	//
	// d := new(int) // *int
	// fmt.Println("Dirección de d: ", d)
	// fmt.Println("Valor de d: ", *d)
	// d = b
	// fmt.Println("Dirección de d: ", d)
	// fmt.Println("Valor de d: ", *d)
	// fmt.Println("Valor de a: ", a)
	// fmt.Println("Valor de b: ", *b)
	// fmt.Println("Valor de c: ", *c)
	//
	// //Usar punteros en funciones

	numero := 2
	fmt.Println("Numero antes de incrementar: ", numero)
	incrementar(&numero)
	fmt.Println("Numero despues de incrementar: ", numero)

}
