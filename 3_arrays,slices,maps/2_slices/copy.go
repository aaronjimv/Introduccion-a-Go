package main

import "fmt"

func main() {

	// copy copia el minimo de elementos de ambos arreglos
	// copy(destino,fuente)

	slice := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(slice), cap(slice)*2)

	fmt.Println(slice)
	fmt.Println(copia)
	fmt.Println("copiando...")

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)

}
