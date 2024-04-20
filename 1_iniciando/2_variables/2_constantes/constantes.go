package main

import "fmt"

// las contastantes no cambian de valor
// a diferencia de las varibles estan no requieren ser usadas.

const (
	OS = "linux"
) 

const (
	Jan = iota + 1 // para incremetar de forma constante, comienza en cero (0)
	Feb
	Mar
	Abr
	May
	Jun
)

func main() {
	fmt.Println(OS)
	fmt.Println(Abr)
}
