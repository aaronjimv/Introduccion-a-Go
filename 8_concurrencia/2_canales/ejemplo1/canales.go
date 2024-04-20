/*
	los canales son para comunicar ejecuciones de codigo concurente via un canal


	nota:
	cuando se envia informacion la info esta del lado derecho y el canal del izquierdo
	cuando se recive info el canal esta en el lado izquierdo y la variable que recibe la info en el derecho

	nota 2:
	el ciclo infinito se interumpe porque esas lineas de codigo se detienen para que el canal envie
	la info al mensaje, esta la recibe y luego finaliza la ejecucion

	nota 3:
	aunque la var nombre esta declarada dentro del for el canal conserva su informacion

*/

package main

import "fmt"

func main() {

	// los canales se crean con make(chan tipo_de_Dato)
	canal := make(chan string)

	go func(canal chan string) {

		for {
			var nombre string
			fmt.Scanln(&nombre)
			// para pasar informaciona una canal se usa: canal <- dato
			canal <- nombre
		}

	}(canal)

	// se tiene que crear un mensaje que reciba la info del canal
	msg := <-canal
	fmt.Println("estoy impriendo lo que recibi del canal: " + msg)
}
