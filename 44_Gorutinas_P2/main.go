package main

import (
	"fmt"
	"time"
)

func main() {

	go animacion(100 * time.Millisecond)
	const n = 45
	resultado := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, resultado)
}

func animacion(retraso time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
