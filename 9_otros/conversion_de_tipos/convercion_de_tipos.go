package main

import (
	"fmt"
	"strconv" //paquete para convertir tipos, en go da error cuando no utilizas una variable o paquete
)

func main() {
	/*
		edad := 21
		fmt.Println("mi edad es: " + edad)

		ERROR: cannot use "mi edad es: " (type untyped string) as type int

		en go no se pueden convertir tipos concatenandolos
	*/

	edad := 21
	edad_str := strconv.Itoa(edad) // Itoa es para int a string

	fmt.Println("mi edad es: " + edad_str)
	// tambien
	fmt.Println("mi edad es: " + strconv.Itoa(edad))

	/*
		edad2 := "40"
		edad2_int := strconv.Atoi(edad2) // Atoi es para string a int

		fmt.Println(edad2_int + 10)

		ERROR: assignment mismatch: 1 variable but strconv.Atoi returns 2 values

		las funciones en go pueden retornar varios valores, Atoi espera dos
	*/

	fmt.Println()

	edad2 := "40"
	edad2_int, _ := strconv.Atoi(edad2) // Atoi es para string a int, con ,_ no da error

	fmt.Println(edad2_int + 10)

}
