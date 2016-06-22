package main

import (
	"fmt"
)

//Type

//Dwclarar nuestro propio tipo

//Dinero Tipo declarado por nosotros
type Dinero int

//Declaramos el metodo String para el tipo Dinero
func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d)
}

func main() {
	var sueldo Dinero
	sueldo = 25000
	fmt.Println("Sueldo: ", sueldo)

	aumento := 10000
	sueldo += Dinero(aumento)

	fmt.Println("Sueldo + Aumento: ", sueldo)

}
