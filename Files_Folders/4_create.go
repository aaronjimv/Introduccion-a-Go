/*
	Para crear un archivo, use la función os.Create.
	Toma el nombre del archivo, lo crea en el directorio de 
	trabajo actual y devuelve un archivo os.y posiblemente un
	error (si por alguna razón no pudo crearlo). 
*/
package main
	
import (
	"os"
)

func main() {
	file, err := os.Create("test_2")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("test")
}
