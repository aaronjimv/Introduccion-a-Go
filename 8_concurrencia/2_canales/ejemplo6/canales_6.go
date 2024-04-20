package main

import (
	"fmt"
	"os"
	"time"
)

func leerEntrada(out chan<- []byte) {
	for {
		datos := make([]byte, 1024)
		n, _ := os.Stdin.Read(datos)
		if n > 0 {
			out <- datos
		}
	}
}

func main() {
	done := time.After(20 * time.Second) // retorna un canal de tipo tiempo (20s)
	eco := make(chan []byte)
	go leerEntrada(eco)
	for {
		select {
		case datos := <-eco:
			os.Stdout.Write(datos)
		case <-done:
			fmt.Println("Se termino el tiempo")
			os.Exit(0)
		}
	}
}

/*
	selec es el swich de los canales
	en vez de evaluar casos evalua cuando en un canal vienen datos
*/
