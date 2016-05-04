package main

import (
	"fmt"
)

func main() {

	//Estructuras de Control: Switch

	//Imprimir el nombre del dia de la semana que el usuario digita

	var dia int

	fmt.Println("Digita el numero del dia de la semana")

	fmt.Scanln(&dia)

	// if dia == 1 {
	// 	fmt.Println("Usted digitó Lunes")
	// } else if dia == 2 {
	// 	fmt.Println("Usted digitó Martes")
	// } else if dia == 3 {
	// 	fmt.Println("Usted digitó Miercoles")
	// } else if dia == 4 {
	// 	fmt.Println("Usted digitó Jueves")
	// } else if dia == 5 {
	// 	fmt.Println("Usted digitó Viernes")
	// } else if dia == 6 {
	// 	fmt.Println("Usted digitó Sabado")
	// } else if dia == 7 {
	// 	fmt.Println("Usted digitó Domingo")
	// } else {
	// 	fmt.Println("Digitó un numero invalido")
	// }
	//
	// fmt.Println("********Fin del Programa************")

	switch dia {
	case 1:
		fmt.Println("Usted digitó Lunes")
		break
	case 2:
		fmt.Println("Usted digitó Martes")
		break
	case 3:
		fmt.Println("Usted digitó Miercoles")
		break
	case 4:
		fmt.Println("Usted digitó Jueves")
		break
	case 5:
		fmt.Println("Usted digitó Viernes")
		break
	case 6:
		fmt.Println("Usted digitó Sabado")
		break
	case 7:
		fmt.Println("Usted digitó Domingo")
		break
	default:
		fmt.Println("Digitó un numero invalido")
		break
	}

	fmt.Println("********Fin del Programa************")

}
