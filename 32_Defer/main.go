package main

import (
	"fmt"
	//"os"
)

//Defer

func primera() {
	fmt.Println("Primera")
}
func segunda() {
	fmt.Println("Segunda")
}

func main() {

	defer primera()
	defer segunda()

	// //Abrimos el archivo
	// f, err := os.Open("texto.txt")
	// //Verificamos que no haya ocurrido ningún error
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// //Creamos un Slice para almacenar los que leemos del archivo
	// data := make([]byte, 175)
	// c, err := f.Read(data)
	// //Verificamos que no haya ocurrido ningún error
	// if err != nil {
	// 	panic(err)
	// }
	// //Imprimimos lo leído
	// fmt.Printf("Cantidad de byte leidos: %d\n Texto leido: \n%q \nerror: %v", c, data, err)
	//
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}

}
