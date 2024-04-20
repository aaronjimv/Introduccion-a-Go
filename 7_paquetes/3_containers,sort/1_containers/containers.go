/*
Además de listas y mapas, Go tiene varias colecciones
más disponibles debajo del paquete container.

El paquete container/list implementa una lista
doblemente vinculada.

	nodo1 -> nodo2 -> nodo3
		1		2		3

Cada nodo de la lista contiene un valor (1, 2 o 3, en este caso)
y un puntero al siguiente nodo. Como se trata de una lista
doblemente vinculada, cada nodo también tendrá punteros al
nodo anterior.
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {

	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println(x)
}

/*
	El valor cero para una Lista es una lista vacía
	(también se puede crear una * Lista usando list.New).

	Los valores se agregan a la lista mediante PushBack.

	Hacemos un bucle sobre cada elemento en
	la lista obteniendo el primer elemento y siguiendo
	todos los enlaces hasta llegar a cero.
*/
