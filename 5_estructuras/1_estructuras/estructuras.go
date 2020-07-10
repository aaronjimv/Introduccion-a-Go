package main 

import "fmt"

//para declarar estructuras se usan las palabras clva type nombre struct
type usuario struct{
	//aqui se colocan las propiedades
	edad int
	nombre,apellido string
}

func main() {

	//para instanciar una escructura, se crea una variable de su tipo
	// var aaron usuario
	// fmt.Println(aaron) // va a imprimir los valores cero de las propiedades => {0  }

	//para declarar e inicializar se usa:
	// aaron := usuario{edad: 21, nombre: "Aaron", apellido: "Jimenez"}
	// fmt.Println(aaron)

	//tambien se pueden declarar por orden pero hay que colocar todos los valores corespondientes
	// aaron := usuario{21,"Aaron","Jimenez"}	
	// fmt.Println(aaron)

	//al utilizar la palabra clave new, se crea un puntero a la estructura
	aaron := new(usuario)

	fmt.Println(aaron) // va imprimir la direccion a la que apuntta

	aaron.nombre = "Aaron" // el .atributo accede al valor en la direcion como en *X
	aaron.apellido = "Jimenez"
	aaron.edad = 21

	fmt.Printf("El usuario se llama %s %s y tiene %d\n", aaron.nombre, aaron.apellido, aaron.edad)
}