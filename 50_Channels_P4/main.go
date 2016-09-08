package main

import (
	"fmt"
	"strconv"
	"time"
)

func enviarMensaje(out chan<- string, numero int) {
	for {

		out <- "Mensaje : " + strconv.Itoa(numero)
		fmt.Println("Enviando mensaje func: ", numero)
	}

}

func imprimir(in <-chan string) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)

	}
}

func main() {
	ch := make(chan string, 5)
	for i := 1; i < 5; i++ {
		go enviarMensaje(ch, i)
	}

	imprimir(ch)
}
