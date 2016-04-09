package main

import "fmt"

func main() {
	var numero int
	multiplicando := 25

	fmt.Println(numero * multiplicando)
	numero = 40
	fmt.Println(numero)
	nombre := "Alejandro"
	nombre = "Pedro"
	nombre, numero = "Rafael", 25
	nombre2 := "Ramon"
	fmt.Println(nombre, nombre2)
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre, nombre2)
	nombre3, nombre := "Maria", "Alejandro"
	fmt.Println(nombre, nombre3)

}
