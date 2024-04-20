package main

import (
	"fmt"
	"time"
)

func main() {
	go animcacio(100 * time.Millisecond)

	const n = 45
	resultado := fib(n)
	fmt.Printf("\rFibinacci(%d) = %d\n", n, resultado)
}

func animcacio(retraso time.Duration) {
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

/*
	como la susecion de fibonacci es con funcion recurciba tarda mucho
	y para dar a entender que el problama no esta paralizado en una
	gorutina se imprime una animacion mientras el programa espera a que se
	calcule el numero, una vez se calcula se imprime y si es fin de main
	por eso la gorutina aunque esta en un for infinita termina y se para

	Entonces...
	mientras se imprima la animacion el programa calcula el numero todo
	al mismo tiempo

	NOTA:
	`\r` significa sobre escribir lo que se imprimio
*/
