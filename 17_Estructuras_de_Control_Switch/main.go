package main

import (
	"fmt"
)

func main() {

	//Estructuras de Control: Switch

	//Imprimir el nombre del dia de la semana que el usuario digita

	// var dia int
	//
	// fmt.Println("Digita el numero del dia de la semana")
	//
	// fmt.Scanln(&dia)
	//
	// // if dia == 1 {
	// // 	fmt.Println("Usted digitó Lunes")
	// // } else if dia == 2 {
	// // 	fmt.Println("Usted digitó Martes")
	// // } else if dia == 3 {
	// // 	fmt.Println("Usted digitó Miercoles")
	// // } else if dia == 4 {
	// // 	fmt.Println("Usted digitó Jueves")
	// // } else if dia == 5 {
	// // 	fmt.Println("Usted digitó Viernes")
	// // } else if dia == 6 {
	// // 	fmt.Println("Usted digitó Sabado")
	// // } else if dia == 7 {
	// // 	fmt.Println("Usted digitó Domingo")
	// // } else {
	// // 	fmt.Println("Digitó un numero invalido")
	// // }
	// //
	// // fmt.Println("********Fin del Programa************")
	//
	// switch dia {
	// case 1:
	// 	fmt.Println("Usted digitó Lunes")
	// case 2:
	// 	fmt.Println("Usted digitó Martes")
	// case 3:
	// 	fmt.Println("Usted digitó Miercoles")
	// case 4:
	// 	fmt.Println("Usted digitó Jueves")
	// case 5:
	// 	fmt.Println("Usted digitó Viernes")
	// case 6:
	// 	fmt.Println("Usted digitó Sabado")
	// case 7:
	// 	fmt.Println("Usted digitó Domingo")
	// default:
	// 	fmt.Println("Digitó un numero invalido")
	// }

	fmt.Println("********Fin del Programa************")

	numero := 26
	switch {
	case numero > 23:
		fmt.Println("Es Mayor que 23")
		fallthrough
	case numero > 25:
		fmt.Println("Es Mayor que 25")
	default:
		fmt.Println("Por lo menos es un Numero XD")
	}

}
