package main

import (
	"fmt"
)

//Panic y Recover

func imprimir() {
	fmt.Println("Hola Alejandro")
	//panic("Error")

	//cadena := recover()

	defer func() {
		cadena := recover()
		fmt.Println(cadena)
	}()
	panic("Error")

}

func main() {

	imprimir()
	fmt.Println("Hola Main")

}
