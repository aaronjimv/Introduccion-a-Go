/*
NOTA 1:
las go rountine se ejecutan en el fondo. entoces cuando llega a la linea de go ---
el programa la ejecuta en el fondo y sigue a la siguiente linea.
SI la otra tambien tine go --- se ejecutara en el fondo y lo que queda seria el fin del
main por lo que el programa se cierra, para evitarlo se coloca un fmt.Scanf() para que el programa
espere a que se precione enter mientras ejecuta las gorountines. Pero eso no es muy practico

	go contador("hola")
	go contador("adios")

	fmt.Scanln()

NOTA 2:
creando un contador con el paquete sync podemos cerrar automaticamente la gorountine cuando llegue
a un numero especifico para eso se pone la gorountine en una funcion anonima

	var wg sync.WaitGroup //crea el contador para parar el gorountine
	wg.Add(1)

	go func () {
		contador("hola")
		wg.Done() // reduce el contador para detener el gorountine
	}()

	wg.Wait() // espera a que llegue a cero
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup //crea el contador para parar el gorountine
	wg.Add(1)

	go func() {
		contador("hola")
		wg.Done() // reduce el contador para parar el gorountine
	}()

	go func() {
		contador("adios")
	}()

	wg.Wait() // espera a que llegue a cero

}

func contador(algo string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, algo)
		time.Sleep(time.Millisecond * 500)
	}
}
