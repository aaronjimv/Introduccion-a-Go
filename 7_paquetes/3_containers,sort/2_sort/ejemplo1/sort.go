/*
	El paquete de clasificación contiene funciones 
	para ordenar datos arbitrarios. Hay varias funciones 
	de clasificación predefinidas (para sectores de ints y flotantes)
*/

package main

import ("fmt" ; "sort")

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {

	kids := []Person{
		{"Jill",9},
		{"Jack",10},
	}

	fmt.Println(kids)

	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

/*
	La función Sort en sort toma un sort.Interface y lo ordena.
	El sort.Interface requiere tres métodos: Len, Less y Swap.

	Len debería devolver la longitud de lo que estamos clasificando.
	Para un segmento, simplemente devuelve len (ps).
	
	Less se utiliza para determinar si el artículo en la posición i es 
	estrictamente menor que el artículo en la posición j.
	En este caso, simplemente comparamos
	ps [i] .Name con ps [j] .Name.
	
	Swap intercambia los artículos.
*/
