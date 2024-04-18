/*
	A menudo, queremos recorrer una carpeta de forma 
	recursiva (leer el contenido de la carpeta, todas
	las subcarpetas, todas las subcarpetas, etc.). 
	Para facilitar esto, hay una función Walk en el 
	paquete path / filepath:
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

/*
	La función que pasa a Walk se llama para cada archivo y 
	carpeta en la carpeta raíz (en este caso, . ). 
	Se han pasado tres argumentos: path, que es la ruta al archivo; 
	info, que es la información del archivo (la misma información 
	que obtiene al usar os.Stat); y err, que es cualquier error que 
	se recibió al recorrer el directorio. La función devuelve un error
	y puede devolver filepath.SkipDir para dejar de caminar
	inmediatamente.
*/