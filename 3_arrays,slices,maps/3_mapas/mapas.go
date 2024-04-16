package main

import "fmt"

/*
	los mapas son como los diccionarios de python se declaran:
		nombre := make(map[tipo_de_llave], tipo_de_valor)
*/

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	m2 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m2)

	m3 := make(map[string]int)

	// insertar o actualizar un elemento en el mapa
	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	// eliminar un elemento: delete(mapa, llave)
	delete(m3, "Answer")
	fmt.Println("The value:", m["Answer"])

	// comprarbar si la llave posee sus dos valores: llave y elemento
	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	v2, ok2 := m2["Google"]
	fmt.Println("The value:", v2, "Present?", ok2)

	/*
		si el elemento esta en el mapa ok es true sino es false.
		si el elemento no esta en el mapa devuelve el elemento cero.
	*/
}
