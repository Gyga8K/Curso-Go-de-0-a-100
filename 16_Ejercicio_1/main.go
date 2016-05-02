package main

import (
	"fmt"
)

func main() {

	//Contador de Números Impares

	encabezado := `
	****************************
	Contador de Numeros Impares
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
	//Declaramos la variable contador para guardar la cantidad de numeros impares
	var contador int
	//Realizamos un bucle tomando como inicio y fin los numeros digitados
	for i := numero1; i <= numero2; i++ {
		//Evaluamos sin el numero es impar
		if i%2 != 0 {
			//Si el numero es impar
			//incrementamos el valor de la variable contador en 1
			contador++
			//Imprimimos el numero impar
			fmt.Printf("%d es un numero impar\n", i)
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
	fmt.Printf("Entre el %d y el %d hay %d numeros impares.", numero1, numero2, contador)

}
