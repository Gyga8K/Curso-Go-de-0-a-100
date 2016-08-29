package main

import (
	"fmt"
)

func main() {
	//Range
	//TODO Crear funciones faltantes

	nombres := []string{
		"Alejandro",
		"Maria",
		"Pedro",
		"Carlos",
	}

	for index, nombre1 := range nombres {
		// Index = 0
		// nombre nombres[index]
		fmt.Printf("El nombre %q esta en el index %d \n", nombre1, index)
	}

	for _, nombre2 := range nombres {
		fmt.Println(nombre2)
	}

	for index1 := range nombres {
		fmt.Println(index1)
	}

	cadena := "Hola gente, yo soy gyga y esto es GygaCode."

	for index2, letra := range cadena {
		fmt.Printf("La letra %q esta en el index %d \n", letra, index2)
	}

}
