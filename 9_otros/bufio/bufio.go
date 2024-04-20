package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		render := bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa tu nombre: ")
		nombre,err := render.ReadString('\n')
		if err != nil{
			fmt.Println(err)
		}else{
			fmt.Println("Hola "+nombre)
		}*/
	render := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre, _ := render.ReadString('\n')
	fmt.Println("Hola " + nombre)

}
