package main 

import(
	"fmt"
	"net/http"
	"io"
)

func main() {
	
	http.HandleFunc("/hola", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "holaaaaaaaa mundo web")
	});

	http.HandleFunc("/", handler) // llama a la funcion handler cuando hay /
	http.ListenAndServe(":8000", nil) // crea un sevcidor virtual en el puerto :8000
}


	/*	de tener dos parametros en ese orden
		http.ResponseWriter es una estructura que define como se va a responder a la peticios
	   	*http.Request es un puntero a la info de la peticion
		w representa un script de datos de escritura
	*/ 
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hay una nueva peticion")
	io.WriteString(w, "hola mundo web")
}