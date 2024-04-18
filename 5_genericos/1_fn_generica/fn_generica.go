package main

import "fmt"

func main() {
	PrintList("Hola mundo", "gophers", "Hello")
	PrintList(1, 2, 3, 4)
}

// el `...any` le permite a la funcion recibir cualquir tipo de parametro
func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}

/*
// asi no se tiene que repetir el mismo codigo:

func PrintString(list ...string) {
	for _, item := range list {
		fmt.Println(item)
	}
}

func PrintInt(list ...int) {
	for _, item := range list {
		fmt.Println(item)
	}
}
*/
