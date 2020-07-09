package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implementa sort.Interface para [] Person basada en
// el campo Edad.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// Hay dos formas de ordenar unslice. Primero, uno puede definir
	// un conjunto de métodos para el tipo slice, como con ByAge, y
	// llamar sort.Sort. En este primer ejemplo usamos esa técnica.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// La otra forma es usar sort.Slice con un Less personalizado
	// que se puede proporcionar como un cierre. En esto
	// caso no se necesitan métodos. (Y si existen, ellos
	// se ignoran.) Aquí volvemos a ordenar en orden inverso: comparar
	// el cierre con ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

}
