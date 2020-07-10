package main 

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

/*
	let’s walk through factorial(2):

	1. Is x == 0? No (x is 2).
	2. Find the factorial of x − 1

		a. Is x == 0? No (x is 1).
		b. Find the factorial of x − 1.

			i. Is x == 0? Yes, return 1.

		c. Return 1 * 1.

	3. Return 2 * 1.

	Closure y la recursividad son poderosas técnicas de programación 
	que forman la base de un paradigma conocido como programación funcional. 
	La mayoría de las personas encontrarán que la programación funcional es más 
	difícil de entender que un enfoque basado en bucles for, if, 
	declaraciones, variables y funciones simples.
*/

func main() {
	fmt.Println(factorial(5))
}

// al fin entendi al recurcividad nojoda!!!!!!!!!!!!