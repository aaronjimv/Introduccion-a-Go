/*
	1) el archivo debe tener el mimo nombre de la carpeta -> ejemplo/ejemplo.go
	al igual que el package 

	2) no se va a crear una funcion main

	3) las fuciones que empiecen con minusculas no se pueden ejecutarn desde otro paquete
		minuscula -> privada
		Mayuscula -> publica

	4) el ejecutable en el pac. main es func main
		pero en los paquetes es la func init, sirve para inicializar los paquetes
		antes que se usen
	
	5) los paquetes tambien pueden contener atributos exportables

*/

package ejemplo

var Hola string

var nombre string

func init() {
	Hola = "hola ;3"
	nombre = "no me pueden exportar porque empieso en minuscula"
}

func HolaMundo() string{
	return Hola + " Mundo desde el paquete ejemplo"
}

func holaMundoDOS() string{
	return "no me pueden exportar"
}