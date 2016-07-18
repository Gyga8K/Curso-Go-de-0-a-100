package main

import (
	"fmt"
)

//Persona ...
type Persona struct {
	nombre string
	email  string
	edad   int
}

//Nombre ...
func (p Persona) Nombre() string {
	return p.nombre
}

//Email ...
func (p Persona) Email() string {
	return p.email
}

//Moderador ...
type Moderador struct {
	Persona
	Foro string
}

//AbrirForo ...
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir un Foro")
}

//CerrarForo ...
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar un Foro")
}

//Administrador ...
type Administrador struct {
	Persona
	Sección string
}

//CrearForo ...
func (a Administrador) CrearForo() {
	fmt.Println("Abrir un Foro")
}

//EliminarForo ...
func (a Administrador) EliminarForo() {
	fmt.Println("Cerrar un Foro")
}

//Presentarse ...
func Presentarse(p Usuario) {
	fmt.Println("Nombre: ", p.Nombre())
	fmt.Println("Email: ", p.Email())
}

// //PresentarseA ...
// func PresentarseA(a Administrador) {
// 	fmt.Println("Nombre: ", a.Nombre())
// 	fmt.Println("Email: ", a.Email())
// }
//
// //PresentarseM ...
// func PresentarseM(m Moderador) {
// 	fmt.Println("Nombre: ", m.Nombre())
// 	fmt.Println("Email: ", m.Email())
// }

//Usuario ...
type Usuario interface {
	Nombre() string
	Email() string
}

func main() {

	alejandro := Persona{"Alejandro", "a@hmail.com", 29}
	Presentarse(alejandro)
	juan := Moderador{Persona{"Juan", "j@hmail.com", 46}, "Juegos"}
	pedro := Administrador{Persona{"Pedro", "p@hmail.com", 25}, "PC"}
	Presentarse(juan)
	Presentarse(pedro)

	var i Usuario
	i = alejandro
	fmt.Println("i: ", i)
	fmt.Println("i: ", i.Email())
	fmt.Println("i: ", i.Nombre())

	i = juan
	fmt.Println("i: ", i)
	fmt.Println("i: ", i.Email())
	fmt.Println("i: ", i)

}
