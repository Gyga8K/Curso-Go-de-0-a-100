package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// //Enteros con signo
	// var entero8 int8   //Todos los enteros de 8-bit con signo (-128 to 127)
	// var entero16 int16 //Todos los enteros de 16-bit con signo (-32768 to 32767)
	var entero32 int32 //Todos los enteros de 32-bit con signo (-2147483648 to 2147483647)
	var entero64 int64 //Todos los enteros de 64-bit con signo (-9223372036854775808 to 9223372036854775807)
	//
	// //Enteros sin signo
	// var uentero8 uint8   //Todos los enteros de 8-bit sin signo (0 to 255)
	// var uentero16 uint16 //Todos los enteros de 16-bit sin signo (0 to 65535)
	// var uentero32 uint32 //Todos los enteros de 32-bit sin signo (0 to 4294967295)
	// var uentero64 uint64 //Todos los enteros de 64-bit sin signo (0 to 18446744073709551615)
	//
	// //Alias
	// var enteroByte byte //alias para uint8
	var enteroRune rune //alias parar int32
	//
	// //Tipos dependiente de la plataforma
	// var enteroUint uint       //32 o 64 bits
	var enteroInt int //32 o 64 bits
	// var enteroUintptr uintptr //Entero sin signo lo suficientemente grande para contener el valor de un puntero.

	//****************************************
	//Conversión entre tipos

	//Suma int32 y int64

	entero32 = 25 //int32
	entero64 = 85 //int64

	//fmt.Println(entero32 + entero64)
	fmt.Println(entero32 + int32(entero64))

	//****************************************
	//Suma int32 y rune

	enteroRune = 46 //rune
	fmt.Println(entero32 + enteroRune)

	//****************************************
	//Suma int32 y int

	enteroInt = 56 //int
	//fmt.Println(entero32 + enteroInt)
	fmt.Println(entero32 + int32(enteroInt))

	fmt.Println(unsafe.Sizeof(enteroInt), unsafe.Sizeof(entero32))

}
