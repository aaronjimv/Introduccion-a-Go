package main 

import(
	"fmt"
	"time"
)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//generamos los numeros
	go func () {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		close(numero)
	}()

	//los elevamos al cuadrado
	go func () {
			for x := range numero{
				cuadrado <- x * x
			}
			close(cuadrado)
	}()

	//imprimimos en main
	for x := range cuadrado{
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
/*
	AL cerrar los caneles se asecguran que no se envie y reciba mas de la cuenta
	con el close() se cierra el canal y con el for range se itere hasta el
	numero de valores del canal
*/