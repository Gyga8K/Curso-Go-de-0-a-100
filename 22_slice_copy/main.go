package main

import (
	"fmt"
)

func main() {

	//copy

	//Origen = Destino
	// origen := []int{1, 2, 3}
	// destino := []int{3, 4, 5}
	// copy(destino, origen)
	// fmt.Println(origen, destino)

	// //Origen > Destino
	// origen2 := []int{1, 2, 3}
	// destino2 := make([]int, 2)
	// copy(destino2, origen2)
	// fmt.Println(origen2, destino2)

	//Origen < Destino
	origen3 := []int{1, 2}
	destino3 := []int{3, 4, 5}
	copy(destino3, origen3)
	fmt.Println(origen3, destino3)

}
