package main 

import(
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
	for{
		for _, r := range`\|/-`{
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {
	if x < 2{
		return x
	}
	return fib(x-1) + fib(x-2)
}

/*
	como la susecion de fibonacci es con funcion recurciba tarda mucho
	y para dar a atender que el problama no esta paralizado en una 
	gorutina se imprime una anicacion miestras el programa espera a que se
	calcule el numeron una vez se calcula se imprime y si el fin de main
	por eso la gorutina aunque esta en un for infinita termina se para

	Entonces...
	mientras se imprimia la animacion el progrmacalculaba el numero todo 
	al mismo tiempo

	NOTA:
	° \r significa sobre escribir lo que se imprimio
 	° ni idea de como calculo el numero :v
*/