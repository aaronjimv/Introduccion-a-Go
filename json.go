package main 

import(
	"net/http"
	"encoding/json"
)

type Curso struct{
	Titulo string `json:"titulo"`  //asi es como se va a ver en el navegador
	NumeroVideos int `json:"numero_videos"`
}

type Cursos []Curso // arreglo de cursos

func main() {
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		
		curso := Cursos{
			Curso{"Curso de Go", 26},
			Curso{"Curso de Go Web", 20},
			Curso{"Curso de Go Api", 30}, // debe terminar en coma -> ,
		}

		json.NewEncoder(w).Encode(curso)
	})

	http.ListenAndServe(":8001", nil)
}

/*
	json.NewEncoder(w) -> le dice que envie el string de datos al navegador

	ES MUY IMPORTANTE QUE LOS ATRIBUTOS DE LA ESTRUCTURA COMIENSEN EN 
	MINUSCULA PARA QUE SE PUEDAN VER.
	CREO QUE ESO LAS HACE PUBLICAS Y EN MINUSCULA PRIVADAS
*/