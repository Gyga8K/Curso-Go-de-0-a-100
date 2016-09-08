package main

//Channels

import (
	"fmt"

	"time"
)

func imprimirPing(c chan string) {
	var contador int
	for {
		//Recibiendo valores a través del canal
		contador++
		fmt.Println(<-c, " ", contador)
		time.Sleep(time.Second * 1)
	}
}

func enviarPing(c chan string) {

	for {
		//Enviando valores a través del canal
		c <- "Ping"
	}
}
func main() {

	// Creamos un canal hola mundoo
	ch := make(chan string)

	//Llamamos las funciones como Gorutines
	go enviarPing(ch)
	go imprimirPing(ch)

	//Escaneamos la entrada de datos para que no finalice la gorutine “main”
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin...")

}
