package main

import (
	"fmt"
)

func main() {

	//Contador de Números Primos

	encabezado := `
	****************************
	Contador de Numeros Primos
	****************************
	`
	//Imprimimos el encabezado
	fmt.Println(encabezado)
	//Solicitamos el Primer Numero
	fmt.Println("Digita el Primer numero")
	//Declaramos una variable para almacenar el numero digitado
	var numero1 int
	//Leemos el numero digitado y lo guardamos en la variable numero1
	fmt.Scanln(&numero1)
	fmt.Println("Digita el Segundo numero")
	//Declaramos una variable para almacenar el numero digitado
	var numero2 int
	//Leemos el numero digitado y lo guardamos en la variable numero2
	fmt.Scanln(&numero2)
	//Declaramos la variable contador para guardar la cantidad de numeros primos
	var contador int
	//Realizamos un bucle tomando como inicio y fin los numeros digitados
	for i := numero1; i <= numero2; i++ {
		//Evaluamos sin el numero es primo
		if i%2 != 0 {
			//Si el numero es primo
			//incrementamos el valor de la variable contador en 1
			contador++
			//Imprimimos el numero primo
			fmt.Printf("%d es un numero primo \n", i)
		}
	}

	encabezado = `
	*********************
	Resultado del Conteo
	*********************
	`
	//Imprimimos el encabezado
	fmt.Println(encabezado)
	//Imprimimos los resultados
	fmt.Printf("Entre el %d y el %d hay %d numeros primos.", numero1, numero2, contador)

}
