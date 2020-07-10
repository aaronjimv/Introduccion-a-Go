package main

import "fmt"

func main() {
	//variables

	//var nombre_de_la_variables tipo_de_dato
	var x int
	x = 10

	//en una sola linea y asignando tipo de forma automatica seria:
	y := 20

	fmt.Println(x,y)

	nombre := "aaron"
	// con el operador := no se puede reaccinar el valor
	//nombre := "sebastian"
	nombre = "aaron j"

	fmt.Println(nombre)

	// el valor por defecto en bool es false
	var pregunta bool

	fmt.Println(pregunta)
}