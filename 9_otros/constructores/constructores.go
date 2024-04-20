/*
	EN GO NO HAY CONSTRCTORES

	pero por lo que he leido se puede simular modificando el valor cero
	por defecto de los parametros de una estructura

	NOTA:
	como que no hay difenrecia entre crear una instacia atravez de un metodo y
	inicializarla al declalarla

	NOTA 2:
	al parecer puedes hacer algo especifico al momento de crear una estructura

*/

package main

import "fmt"

type Cosa struct {
	nombre string
	numero int
}

func NuevaCosa(algunParametro string, algunNumero int) *Cosa {
	c := new(Cosa)
	c.nombre = algunParametro
	c.numero = algunNumero
	fmt.Println("se ha creado una nueva cosa: " + c.nombre)
	return c
}

func main() {
	caja := NuevaCosa("Caja", 20)

	fmt.Println(caja)
	fmt.Println(*caja)
	fmt.Println(caja.nombre)
	fmt.Println(caja.numero)

	fmt.Println()

	circulo := Cosa{"radio", 4}
	fmt.Println(circulo)
	fmt.Println(circulo.nombre)
	fmt.Println(circulo.numero)
}
