package main

import (
	"fmt"
)

func main() {

	//Operadores Aritmeticos
	// + Suma
	// - Resta
	// * Multiplicacion
	// / Division
	// % Resto

	a := 5
	b := 3

	fmt.Println("5 + 3 = ", a+b)
	fmt.Println("5 - 3 = ", a-b)
	fmt.Println("5 * 3 = ", a*b)
	fmt.Println("5 / 3 = ", a/b)
	fmt.Println("5 % 3 = ", a%b)

	// Operadores de Comparación
	// == Igual que
	// != Diferente que
	// < Menor que
	// <= Menor o Igual que
	// > Mayor que
	// >= Mayor o Igual que
	c := 3
	fmt.Println("******************************************")

	fmt.Println("5 == 3 = ", a == b)
	fmt.Println("3 == 3 = ", b == c)
	fmt.Println("5 != 3 = ", a != b)
	fmt.Println("3 != 3 = ", c != b)
	fmt.Println("5 < 3 = ", a < b)
	fmt.Println("5 <= 3 = ", a <= b)
	fmt.Println("3 <= 3 = ", c <= b)
	fmt.Println("5 > 3 = ", a > b)
	fmt.Println("5 >= 3 = ", a >= b)
	fmt.Println("3 >= 3 = ", c >= b)

	// Operadores de asignación
	// += || a += b || a = a + b
	// -= || a -= b || a = a - b
	// *= || a *= b || a = a * b
	// /= || a /= b || a = a / b
	// %= || a %= b || a = a % b

	// Operadores Logicos
	// && AND
	// || OR
	// ! Negacion
	fmt.Println("******************************************")
	fmt.Println("AND - &&")
	fmt.Println("true && true = ", true && true)
	fmt.Println("false && true = ", false && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("false && false = ", false && false)
	fmt.Println("OR - ||")
	fmt.Println("true || true = ", true || true)
	fmt.Println("false || true = ", false || true)
	fmt.Println("true || false = ", true || false)
	fmt.Println("false || false = ", false && false)
	fmt.Println("Negacion - !")
	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false)
	fmt.Println("******************************************")

	// Jerarquía de Operadores
	// ()
	// * / %
	// + -
	// == != < <= > >=
	// &&
	// ||
	fmt.Println((5 - 3) * 5)

	//Operadores de incremendo y decremento
	//++ || a++ || a = a + 1
	//-- || a-- || a = a - 1
	//b + a++ - 3 + 5
	//b + ++a - 3 + 6
	a++
	fmt.Println("a++ ", a)

}
