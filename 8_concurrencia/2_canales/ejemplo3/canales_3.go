package main

import (
	"fmt"
	"time"
)

func imprimirPing(c chan string) {
	var cantador int
	for {
		//recibiendo valores a traves del canal
		cantador++
		fmt.Println(<-c, " ", cantador)
		time.Sleep(time.Second * 1)
	}
}

func enviarPing(c chan string) {
	for {
		// enviamos valores a traves del canal
		c <- "Ping"
	}
}

func main() {

	// creamos un canal
	ch := make(chan string)

	// llamamos las fuciones como goroutine
	go enviarPing(ch)
	go imprimirPing(ch)

	// esperamos a que lea algo para finalizar el main y parar las goroutine
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin...")
}

/*
	aunque ambas goroutine estan en for infinitas, automaticamente se
	paran para recibir y enviar informacion a traves de los canales.

	Primero el ciclo infinito se para para enviar la informacion
	Luego el otro ciclo infinito recibe la informacion y luego se vuelve a
	parar para recibir otra vez

	El programa se para cuando se escanea algo para cerrar el main
*/
