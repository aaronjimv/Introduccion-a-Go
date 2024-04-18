package main

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func main() {
	mario := Person{
		Name:        "Mario",
		Age:         42,
		HasChildren: false,
	}

	luigi := Person{"Luigi", 33, true}

	wario := Person{Age: 32} // como no se declaro el resto, obtienen sus valores por defecto.

	fmt.Printf("%+v\n", mario)
	fmt.Printf("%+v\n", mario.Name) // para acceder a un unico valor

	fmt.Printf("%+v\n", luigi)
	fmt.Printf("%+v\n", luigi.HasChildren)

	fmt.Printf("%+v\n", wario)
	fmt.Printf("%+v\n", wario.Age)
	fmt.Println()

	// original
	fmt.Println("Valor original: ")
	fmt.Printf("%+v\n", luigi)

	// Usando punteros
	//
	walugi := &luigi // crea un punturo a 'luigi'
	walugi.Age = 60  // modiifca el valor de 'Age' y como es el puntero modifica tambien el valor 'Age' de luigi

	fmt.Println("Punteros: ")
	fmt.Printf("%+v\n", luigi)
	fmt.Printf("%+v\n", walugi)
}
