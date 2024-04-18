package main

import "fmt"

// los campos anonimos simulan la herencia en poo
type Humano struct {
	nombre string
}

func (this Humano) Hablar() string {
	return "bla bla bla..."
}

type Tutor struct {
	Humano // Tutor tine un campo de tipo Humano
}

func (this Tutor) Hablar() string {
	return this.Humano.Hablar() + " holaaaaaaaaaaaaaaaaaa"
}

func main() {

	tutor := Tutor{Humano{"gophers"}}

	fmt.Println(tutor.nombre)
	// aunque Tutor no tiene un campo nombre, Humano	si
	// tambien sirve tutor.Humano.nombre pero es cuando hay dos
	// campos anonimos con el mismo nombre de atributo

	//tambien se pueden acceder a los metodos de la estructura Humano
	fmt.Println(tutor.Hablar())

	// tambien se pueden sobreescribir los metodos

}
