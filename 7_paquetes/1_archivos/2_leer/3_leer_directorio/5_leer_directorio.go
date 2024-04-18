/*
	Para obtener el contenido de un directorio, usamos la misma 
	función os.Open pero le damos un ruta de directorio en lugar 
	de un nombre de archivo. Luego llamamos al método Readdir
*/
package main

import (
	"fmt"
	"os"
)
func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
/*
	Readdir toma un solo argumento que limita el número de 
	entradas devueltas. Al pasar -1, devolvemos todas 
	las entradas.
*/