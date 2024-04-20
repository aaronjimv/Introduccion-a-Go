package main

import "fmt"

// arreglos literal
// son un poco mas flexible ya que no hay que declarar el tamaño con anterioridad
// se usa:nombre_arrglo := [...] tipo {valores}

func main() {
	// var flags [3]string
	// flags[0] = "🎏"
	// flags[1] = "🏴"
	// flags[2] = "🏳️"

	flags := [...]string{"🎏", "🏴", "🏳️", "🏴‍☠️", "🏴"}

	// el tamaño despues de declarar, no se puede cambiear
	// flags[6] = "🏁"

	fmt.Println(flags)
}
