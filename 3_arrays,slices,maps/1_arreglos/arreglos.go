package main 

import "fmt"

func main() {
	
	//los arreglos en go son estaticos, se tiene que declarar el tamaño y tiene que ser el mismo tipo de dato
	var arreglo [10] int

	fmt.Println(arreglo)

	// al hacerlo de esta forma hay que colocar {} adentro se puede poner el valor
	arreglo2 := [3] int {}

	fmt.Println(arreglo2)

	//para la longitub de usa len()
	fmt.Println(len(arreglo))

	for i:=0;i<len(arreglo);i++{
		arreglo[i]=i
	}
	for i:=0;i<len(arreglo);i++{
		fmt.Println(arreglo[i])
	}

	//las matrices son: var nombre[cantidad de arreglos][numero de valores]
	var matriz[3][2] int
	fmt.Println(matriz)

}