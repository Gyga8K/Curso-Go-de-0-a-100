package main

import (
	"fmt"
	"time"
)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros
	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		close(numero)
	}()

	//Lo elevamos al cuadrado
	go func() {
		for {
			x, ok := <-numero
			if !ok {
				break
			}
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	//Imprimimos en main
	for {
		x, ok := <-cuadrado
		if !ok {
			break
		}
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}
