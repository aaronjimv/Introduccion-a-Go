package main

import "fmt"

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func main() {
	product1 := Product[uint]{ID: 1, Description: "Zapatos", Price: 30}
	product2 := Product[string]{ID: "6f9b12c3-d9e8-4cfc-b841-b002e2fafd41", Description: "Camisa", Price: 40}

	fmt.Println(product1)
	fmt.Println(product2)
}

/*
// la funciones de los genericos es no tener que repetir el codigo
// cuando el unico cambio es el tipo de dato.

type Product(unit) struct {
	ID          unit
	Description string
	Price       float64
}

type ProductVer2(string) struct {
	ID          string
	Description string
	Price       float64
}

// Con los genericos se puede utilizar una sola estructura
// de la siguiente forma:

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}
*/
