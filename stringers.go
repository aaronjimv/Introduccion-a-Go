package main 

import "fmt"

// el metoso: String () string 
// es como el ToString() en python

type  Persona struct{
	nombre string
	edad int
}

func (this Persona) String() string {
	// tiene que regresar el metodo Sprintf()
	return fmt.Sprintf("%v tiene %v" , this.nombre, this.edad)
}


func main() {
	
	aaron := Persona{"Aaron", 21}
	sebastian := Persona{"Sebastian", 21}

	fmt.Println(aaron, sebastian)	
}