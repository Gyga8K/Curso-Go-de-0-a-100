package main

import (
	"fmt"
)

//Punto contiene los valores x, y
type Punto struct {
	x, y int
}

//Punto3D contiene los valores x,y,z
type Punto3D struct {
	x, y, z int
	*Punto3D
}

// //
// // //OpPunto tiene los metodos para realizar operaciones con puntos
// type OpPunto struct {
// }

//
func main() {
	p := Punto{}
	fmt.Println("p: ", p)

	p2 := Punto3D{
		5,
		6,
		4,
		&Punto3D{
			6,
			4,
			6,
			nil,
		},
	}
	fmt.Println("p2: ", p2)

	//Comparando Estructuras

	a := Punto{5, 6}
	b := Punto{7, 4}
	fmt.Println("a == b : ", a == b)
	//
	c := Punto{7, 4}
	fmt.Println("b == c :", b == c)

	figuras := make(map[Punto]string)
	figuras[a] = "Hola Mundo"
	fmt.Println("figuras[a] :", figuras[a])

}
