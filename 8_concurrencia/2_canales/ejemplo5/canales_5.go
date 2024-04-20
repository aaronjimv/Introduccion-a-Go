package main

import (
	"fmt"
	"time"
)

func generarNumeros(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func elevarAlCuadrado(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func imprimir(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//generamos los numeros
	go generarNumeros(numero)
	//los elevamos al cuadrado
	go elevarAlCuadrado(numero, cuadrado)
	//imprimimos en main
	imprimir(cuadrado)
}

/*
	con los canales de una sola direccion de asegura que un canal
	solamente pueda enviar o recivir datos

	out chan<- int
	in <-chan int
*/
