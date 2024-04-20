package main

import "fmt"

func main() {
	//variables
	// en go todas las variables se tienen que usar.

	//var nombre_de_la_variables tipo_de_dato
	var x int
	x = 10

	//en una sola linea y asignando tipo de forma automatica seria:
	y := 20

	fmt.Println(x,y)

	nombre := "Mario"
	fmt.Println(nombre)

	// con el operador := no se puede reasignar el valor
	//nombre := "Luigi"
	nombre = "Luigi"

	fmt.Println(nombre)

	// el valor por defecto en bool es false
	var pregunta bool

	fmt.Println(pregunta)
}