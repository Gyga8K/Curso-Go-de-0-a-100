package main

import (
	"fmt"
	"os"
	"time"
)

func leerEntrada(out chan<- []byte) {

	for {
		datos := make([]byte, 1024)
		n, _ := os.Stdin.Read(datos)
		if n > 0 {
			out <- datos
		}
	}
}

func main() {
	done := time.After(20 * time.Second)
	eco := make(chan []byte)
	go leerEntrada(eco)
	for {
		select {
		case datos := <-eco:
			os.Stdout.Write(datos)
		case <-done:
			fmt.Println("Se terminó el tiempo")
			os.Exit(0)
		}
	}

}
