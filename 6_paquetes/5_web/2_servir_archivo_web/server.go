/*
	la forma mas facil de servir archivos en go es con
	http.ServeFile(w,r,"archivo")

	w: string de datos
	r: la peticion
	archivo: nombre el archivlo, debe ser en la misma carpeta donde esta
		el server.go 
*/

package main 

import(
	"net/http"
	"fmt"
)

func main() {
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		
		//http.ServeFile(w,r,"index.html") -> sirve el archivo index.html
		
		fmt.Println(r.URL.Path)
		http.ServeFile(w,r,r.URL.Path[1:]) // -> sirve todos los archivos dentro de la carpeta

	})
	
	http.ListenAndServe(":8000", nil) // crea un sevcidor virtual en el puerto :8000
}

/*
	http.ServeFile(w,r,r.URL.Path[1:])

	r.URL -> devuelve la url
	r.URL.Path -> devuelve todo lo que este despues de la / , inclueida la /
	r.URL.Path[1:] -> omite el caracter cero que es la / y todo lo demas es
						la ruta donde tiene que buscar el archivo

*/