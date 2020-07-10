/*
	lo malo de la forma anterior es que expone todos los archivos y 
	puede ser peligroso
	
	con fileServe := http.FileServer(http.Dir("public"))
	se acegura que solo se pueda acceder a los archivos dentro de la carpeta public
*/

package main 

import(
	"net/http"
)

func main() {
	
	fileServe := http.FileServer(http.Dir("public"))

	http.Handle("/", http.StripPrefix("/", fileServe))
	http.ListenAndServe(":8000", nil) 
}
