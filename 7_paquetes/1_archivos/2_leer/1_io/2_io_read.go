package main 


/*
	esta forma de leer archivo no es recomendado para archivos muy grandes
*/


import(
	"fmt"
	"io/ioutil"
)

func main() {
	
	// lee el archivo, el ./ significa en al misma carpeta
	// retorna dos variable: la info del archivo y un error
	archivo, err := ioutil.ReadFile("./hola")	

	if err != nil{
		fmt.Println("hubo un error")
	}

	// la info esta en bytes por eso hay que transformarlo en un string
	str := string(archivo)
	fmt.Println(str)

}