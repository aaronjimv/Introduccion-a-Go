package main 

/*
	go no tiene ecepiones pero los errors son parecidos, la diferencia es 
	que error es un tipo de dato
*/

import(
	"fmt"
	"math"
	"errors"
)

func main() {
	
	resultado, err := raiz(16)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(resultado)
	}
}


func raiz(x float64) (float64, error) {
	if x < 0{
		// regresa cero y un nuevo error
		return 0, errors.New("No acepta numeros negativos")
	}

	// regresara la raiz cuadrada y error nulo
	return math.Sqrt(x), nil
}