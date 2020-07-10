package main 

import "fmt"

type usuario struct{

	edad int
	nombre,apellido string
}

// go no tinene clases asi que los metodos son de las escructuras
// para agregar un metodo a una estructura es:
// func (id_dentro_de_la_funcion nombre_estructura) nombre_funcion() tipo_dato_que_retorna

// TRUCO: EL ID PUEDE LLAMARCE this Y ASI PARECE POO
func (this usuario) imprimir() string{
	return this.nombre + " " + this.apellido
}

// los parametors que se le pasan a la funcion de la estructura es una copia
// por lo que crea otra estructura aparte para solucionarlo se pasa un puntero
// y asi se puede modificar las propiedades

func (this *usuario) set_nombre(n string) {
	this.nombre = n
}

func main() {

	aaron := new(usuario)

	aaron.nombre = "Aaron" 
	aaron.apellido = "Jimenez"
	aaron.edad = 21

	fmt.Println(aaron.imprimir())

	fmt.Println()

	aaron.set_nombre("Sebastian") // como es un puntero si modifica el nombre

	fmt.Println(aaron.imprimir())
}