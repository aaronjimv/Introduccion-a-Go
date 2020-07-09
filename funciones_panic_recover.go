package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC!!!!!!")
}

/*
	Un panic generalmente indica un error del programador 
	(por ejemplo, intentar acceder a un índice de una matriz que
	está fuera de los límites, olvidarse de inicializar un mapa, etc.)
	o una condición excepcional de la que no hay una manera fácil 
	de recuperarse (de ahí el nombre de pánico) .

	recover detiene el panic y devuelve el valor que se pasó 
	a la llamada al panic.
*/