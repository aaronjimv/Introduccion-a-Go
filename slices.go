package main 

import "fmt"

func main() {
	
	//los slices son areglos pero dimanamicos se pueden declarar vacios
	//  var slice [] int 
	//con solo colocarlo vacio es un slice, para agregarle valores se usan otros metodos

	slice2 := [] int {1,2,3,4,5}
	fmt.Println(slice2)

	//los slices vacios tinen el valor nulo, NO CERO
	//para saber si esta vacio se usa:
	// var slice3 [] int //vacio
	slice3 := [] int {1,2,3}
	if slice3 == nil{
		fmt.Println("es nulo")
	}else{
		fmt.Println(slice3)
	}

	//los slices tienen 3 datos:
	// la longitub
	// un punturo al arreglo que esta referenciando
	// la capacidad

	// para combertir un arreglo en un slice es:
	arreglo := [] int {1,2,3,4}
	slice4 := arreglo[:4] // el :4 es slicing como en ṕython
	fmt.Println(slice4)
}