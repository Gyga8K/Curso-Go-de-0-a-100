package main

import (
	"fmt"
)

func main() {
	//Range

	nombres := []string{
		"Alejandro",
		"Maria",
		"Pedro",
		"Carlos",
	}

	for index, nombre := range nombres {
		// Index = 0
		// nombre nombres[index]
		fmt.Printf("El nombre %q esta en el index %d \n", nombre, index)
	}

	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	for index := range nombres {
		fmt.Println(index)
	}

	cadena := "Hola gente, yo soy gyga y esto es GygaCode."

	for index, letra := range cadena {
		fmt.Printf("La letra %q esta en el index %d \n", letra, index)
	}

}
