package main

import "fmt"

var razaGoku = "Saiyajin"

func main() {
	//Nombrar Variables
	numero := 25
	_nombre := "Alejandro"
	numero2 := 54
	nombreUsuario := "admin"
	Numero := 46
	fmt.Println(numero, numero2, _nombre,
		nombreUsuario, Numero)
	var (
		dios               = "Goku"
		enemigo1, enemigo2 = "Babidi", "Cell"
	)
	var numero3 = 66
	fmt.Println(dios, enemigo1, enemigo2, numero3)

	//Scope

	fmt.Println("La raza de Goku es: " + razaGoku)
	imprimir()

}

func imprimir() {
	fmt.Println("La raza de Goku es: " + razaGoku)
}
