package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	var condicion int

	for true {
		fmt.Println("bienvenido al menu para continuar oprima 1 para salir oprima 2")
		fmt.Scanf("%d\n", &condicion)

		if condicion == 1 {
			fmt.Println("continuando...")
		} else if condicion == 2 {
			break
		} else {
			fmt.Println("OPCION INVALIDA")
		}
	}

}
