package main 

import "fmt"

func main() {

	// se pueden crear funciones dentro de otras fuciones
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,1))
	// add es una variable focal de tipo func(int, int)

	//cuando se crean una funcion local como esta, 
	//tambien tiene acceso a las variables locales
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// agrea 1 a la variable local de main x

	/*
	Una función como esta junto con las variables no locales, de la funcion, 
	a las que hace referencia se conoce como closure. 
	En este caso, increment y la variable x forman el closure.

	Una forma de usar closure es escribiendo una función que devuelva 
	otra función, que cuando se llama, puede generar una secuencia
	de números.
	*/

	pares := generadorPares()
	fmt.Println(pares()) // 0
	fmt.Println(pares()) // 2
	fmt.Println(pares()) // 4

}

func generadorPares() func() int {
	i := 0
	return func() (ret int) { // va a retornar la variable ret
		ret = i
		i += 2
		return
	}
}
