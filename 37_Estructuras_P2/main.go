package main

import (
	"fmt"
)

//Persona contiene los campos nombre y apellido
type Persona struct {
	Nombre   string
	Apellido string
}

//Estudiante contiene el campo persona y carrera
type Estudiante struct {
	Persona
	Carrera string
}

//Profesor tiene los campos Estudiante y Carrera
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	//Declaramos una variable del tipo Estudiante
	alejandro := Estudiante{
		Persona{
			Nombre:   "Alejandro",
			Apellido: "Arnaud",
		},
		"Informatica",
	}
	fmt.Println("Alejandro: ", alejandro)

	//Accediendo a los campos
	fmt.Println("Nombre: ", alejandro.Nombre)
	fmt.Println("Apellido: ", alejandro.Apellido)
	fmt.Println("Carrera: ", alejandro.Carrera)

	//Declaramos una clase de tipo Profesor
	pedro := Profesor{
		Estudiante{
			Persona{
				"Pedro",
				"Almonte",
			},
			"Cotabilidad",
		},
		"Informatica",
	}
	fmt.Println("Pedro: ", pedro)
	//
	//Accediendo a los campos
	fmt.Println("Nombre: ", pedro.Nombre)
	fmt.Println("Apellido: ", pedro.Apellido)
	fmt.Println("Carrera: ", pedro.Carrera)
	//
	fmt.Println("Carrera Estudiantes: ", pedro.Estudiante.Carrera)
	//
	//Declarando una variable de tipo Profesor
	manuel := Profesor{Estudiante{Persona{"Manuel", "Peralta"}, "Ing. Industrial"}, "Informatica"}
	fmt.Println("Manuel: ", manuel)
	//
	var jose Profesor
	jose.Nombre = "José"
	jose.Apellido = "Contreras"
	jose.Estudiante.Carrera = "Educación"
	jose.Carrera = "Mercadeo"
	fmt.Println("José: ", jose)

}
