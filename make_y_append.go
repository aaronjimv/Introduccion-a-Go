package main 

import "fmt"

func main() {
	
	//la palabra clave make clea un slice
	// su forma es []tipo_de_dato,longitub,capacidad
	slice := make([]int,3,6)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // cap() es para ver la capacidad

	//con append se agrega un nuevo elemento al slice
	slice = append(slice,1,2,3,4,5,6,7,8) 
	//al parecer se puede revasar la capacidad
	fmt.Println(slice)
}