/*
	lo malo de la forma anterior es que expone todos los archivos y
	puede ser peligroso

	con fileServe := http.FileServer(http.Dir("public"))
	se asegura que solo se pueda acceder a los archivos dentro de la carpeta public
*/

package main

import (
	"net/http"
)

func main() {

	fileServe := http.FileServer(http.Dir("public"))

	http.Handle("/", http.StripPrefix("/", fileServe))
	http.ListenAndServe(":8000", nil)
}

// Nota:
// debes moverte en la terminal hasta este directorio
// cd 7_paquetes/5_web/2_servir_archivo_web/ejemplo2/
