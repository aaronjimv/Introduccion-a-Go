package main

import "fmt"

// para declarar estructuras se usan las palabras clava type nombre struct
type usuario struct {
	//aqui se colocan las propiedades
	edad             int
	nombre, apellido string
}

func main() {

	//para instanciar una escructura, se crea una variable de su tipo:
	//
	// var mario usuario
	// fmt.Println(mario) // va a imprimir los valores cero de las propiedades => {0  }

	//para declarar e inicializar se usa:
	//
	// mario := usuario{edad: 21, nombre: "Mario", apellido: "Bros"}
	// fmt.Println(mario)

	//tambien se pueden declarar por orden pero hay que colocar todos los valores corespondientes
	//
	// mario := usuario{21,"Mario","Bros"}
	// fmt.Println(mario)

	//al utilizar la palabra clave new, se crea un puntero a la estructura
	mario := new(usuario)

	fmt.Println(mario) // va imprimir la direccion a la que apunta

	mario.nombre = "Mario" // el .atributo accede al valor en la direcion como en *X
	mario.apellido = "Bros"
	mario.edad = 21

	fmt.Printf("El usuario se llama %s %s y tiene %d\n", mario.nombre, mario.apellido, mario.edad)
}
