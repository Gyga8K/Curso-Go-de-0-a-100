package main

// Gorutinas

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg es utilizado para indicarle al programa que debe esperar
// que finalicen las gorutinas
var wg sync.WaitGroup

func main() {
	// Añadimos 2 a wg para que espero que finalicen 2 gorutinas.
	wg.Add(2)
	fmt.Println("Iniciamos las gorutinas...")
	//Lanzamos la gorutina con la etiqueta "A"
	go imprimirCantidad("A")
	//Lanzamos la gorutina con la etiqueta"B"
	go imprimirCantidad("B")
	// Esperamos a que las gorutinas finalicen.
	fmt.Println("Esperando que Finalicen...")
	wg.Wait()
	fmt.Println("\nTerminando el programa")
}

func imprimirCantidad(etiqueta string) {
	// Llamamos la funcion Done() de wg para indicale que la gorutina termino.
	defer wg.Done()
	// Espera aleatoria
	for cantidad := 1; cantidad <= 10; cantidad++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Cantidad: %d de %s\n", cantidad, etiqueta)
	}
}
