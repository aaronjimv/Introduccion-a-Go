package main

import "fmt"

func main() {
	fmt.Println(sum(2))
	fmt.Println(sum(2, 3))
	fmt.Println(sum(2, 3, 12))
	fmt.Println(sum(2, 3, 12, 1, 24))

	// esta es una funcion anonima
	func(name string) {
		fmt.Println("✋ Hello", name)
	}("gophers")

}

// es una funcion variadic por que tiene '...' puede recibir n cantidad de parametros
func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
