package main

import "fmt"

// el metodo: String () string
// es como el ToString() en python

type Persona struct {
	nombre string
	edad   int
}

func (this Persona) String() string {
	// tiene que regresar el metodo Sprintf()
	return fmt.Sprintf("%v tiene %v", this.nombre, this.edad)
}

func main() {

	bill := Persona{"Bill", 50}
	elon := Persona{"Elon", 50}

	fmt.Println(bill, elon)
}
