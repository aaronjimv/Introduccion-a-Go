package main

import (
	"fmt"
	"os"
)

/*
defer siempre se ejecuta despues del retorno de la funcion,
en este caso se ejecuta despues que main() retorne, es decir,
despues que termina main()

Por eso es bueno para cerrar procesos, archivos, coneciones, etc.
*/
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello Word!"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
