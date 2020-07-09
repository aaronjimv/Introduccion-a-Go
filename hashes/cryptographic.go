/*
	Las funciones hash criptográficas son similares a sus 
	contrapartes no criptográficas, pero tienen la propiedad 
	adicional de ser difíciles de revertir. Dado el hash 
	criptográfico de un conjunto de datos, es extremadamente 
	difícil determinar qué hizo el hash. Estos hashes a menudo 
	se usan en aplicaciones de seguridad.

	Una función hash criptográfica común se conoce como SHA-1. 
	Así es como se usa:
*/
package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {

	h := sha1.New()

	h.Write([]byte("test"))

	bs := h.Sum([]byte{})

	fmt.Println(bs)
}
/*
	Este ejemplo es muy similar al crc32, porque tanto crc32 
	como sha1 implementan la interfaz hash.Hash. La principal 
	diferencia es que mientras que crc32 calcula un hash de 32 bits,
	sha1 calcula un hash de 160 bits. No hay un tipo nativo para 
	representar un número de 160 bits, por lo que utilizamos un 
	segmento de 20 bytes.
*/
