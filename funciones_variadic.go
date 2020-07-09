package main

import "fmt"


func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1,2,3,6))

	xs :=  []int{1,2,3,6}
	fmt.Println(add(xs...))
}
/*
	usando (...) antes del tipo del parametro, se puede indicar qur toma cero o
	mas de esos parametros

	se invoca como cualquier otra funcion y le podemos pasar cuantos int queramos 

	tambien se puede pasar un slice de int
*/