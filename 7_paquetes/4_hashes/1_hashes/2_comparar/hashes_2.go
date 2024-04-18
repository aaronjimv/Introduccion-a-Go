/*
	Un uso común de crc32 es comparar dos archivos.
	Si el valor de Sum32 para ambos archivos es el mismo, 
	es muy probable (aunque no 100% seguro) que los archivos 
	sean iguales. Si los valores son diferentes, entonces 
	los archivos definitivamente no son los mismos:
*/
package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)

	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()

	// create a hasher
	h := crc32.NewIEEE()

	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, errH := io.Copy(h, f)

	// we don't care about how many bytes were written, but we do 
	// want to handle the error
	if errH != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {

	h1, err := getHash("test1")

	if err != nil {
		return
	}

	h2, err := getHash("test2")

	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}


