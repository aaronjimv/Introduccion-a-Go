package main

import "fmt"

/*
	Go tiene una declaración especial llamada defer que
	programa una llamada de función para que se ejecute
	después de que la función se complete
*/

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}

/*
	A menudo se usa cuando los recursos necesitan ser
	liberados de alguna manera. Por ejemplo,
	cuando abrimos un archivo, debemos asegurarnos
	de cerrarlo más tarde.

		f, _ := os.Open(filename)
		defer f.Close()

	Esto tiene tres ventajas:
	• Mantiene nuestra llamada close cerca de nuestra
	llamada open para que sea más fácil de entender.

	• Si nuestra función tuviera múltiples declaraciones de
	retorno (quizás una en un if y una en un else),
	Close sucederá antes que los dos.

	• Las funciones defer se ejecutan incluso si
	se produce un pánic en tiempo de ejecución.
*/
