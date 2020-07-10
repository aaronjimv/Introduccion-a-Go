package main 

import "fmt"

func main() {
	
	/*
		PUNTEROS
		1. es una direccion de memiria
		2. en lugar del valor tenemos la direccion en donde esta el valor
		3. X,Y => as123d => 5
		4. X => as123d => 6
		5. Y ¿? => 6

		*T => *int, *string...
		valor cero => nil
	*/

		var x,y *int
		entero := 5

		x = &entero // se debe dar valor a los punteros atravez de otra variable
		y = &entero

		fmt.Println(x) //direccion en memoria
		fmt.Println(y)
		fmt.Println(*x)	// valor en la direccion en memoria
		fmt.Println(*y)


		fmt.Println()

		*x = 6 // se modifico el vlor en la direccion en memoria de x

		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(*x)
		fmt.Println(*y) // tambien modifico el de y porque ambos apuntan a la misma  direccion en memoria 
}