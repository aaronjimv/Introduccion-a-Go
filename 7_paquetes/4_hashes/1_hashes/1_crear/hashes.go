/*
	Una función hash toma un conjunto de datos y lo 
	reduce a un tamaño fijo más pequeño. Los hashes se 
	usan con frecuencia en la programación para todo,
	desde buscar datos hasta detectar cambios fácilmente. 
	Las funciones de hash en Go se dividen en dos categorías:
	criptográficas y no criptográficas.

	Las funciones hash no criptográficas se pueden encontrar 
	debajo del paquete hash e incluyen adler32, crc32, crc64 y 
	fnv. 

	Aquí hay un ejemplo usando crc32:
*/
package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}
/*
	El objeto hash crc32 implementa la interfaz de Writer,
	por lo que podemos escribir bytes como cualquier otro Writer.
	Una vez que hemos escrito todo lo que queremos, llamamos 
	a Sum32 () para devolver un uint32.
*/
