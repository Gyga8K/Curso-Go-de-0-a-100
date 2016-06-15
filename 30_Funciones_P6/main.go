package main

import (
	"fmt"
	"strings"
)

func main() {

	cadena := "aeiou"

	cadena = strings.Map(func(r rune) rune {
		return r + 1

	}, cadena)

	fmt.Println(cadena)

}
