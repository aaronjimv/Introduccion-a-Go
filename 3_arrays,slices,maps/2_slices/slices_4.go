package main

import "fmt"

func main() {
	// len(): numero de elementos en el slice
	// cap(): numero de elementos del array origen, a partir del índice donde se creo el slice

	animals := [5]string{"🦍", "🐕", "🦮", "🐦", "🐘"}
	pets := animals[1:3]               // valores: "🐕", "🦮"
	pets = append(pets, "🐈", "🐕", "🐱") // el append() sive para agregar mas valores al slice

	// Al momente de superar la capacidad del slice, go lo dupilca creando un nuevo slice:
	// 		Array[4]{"🐕", "🦮", "🐦", "🐘"}
	// 		new Array[8]{"🐕", "🦮",  "🐈", "🐕", "🐱"}

	// make() => sive para crear un slice, se declara el tipo, tamaño y la capacidad
	//
	// sin make() => pets := []string{"🐕", "🦮"}
	// con make() => pets := make([]string, 0, 3)

	// El valor nulo de los slice es 'nil'
	//
	// pets = append(pets, "🐈", "🐕", "🐱", "🐕")
	// pets := []string{} // pets = 'nil'

	fmt.Println("animals:", animals)
	fmt.Println("pets:", pets)
	fmt.Println("tamaño pets:", len(pets))
	fmt.Println("capacidad pets:", cap(pets))
	fmt.Println("valor cero:", pets == nil)
}
