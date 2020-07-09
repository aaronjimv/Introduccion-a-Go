package main 

import(
	"fmt"
	"time"
	"strings"
)

// goroutines es para concurrencia, que hayan varios procesos en ejecucion
// se pueden ejecutar varias lineas de codigo al mismo tiempo

func main() {
	
	go mi_nombre_lentooooo("aaron") //con solo poner go adelante la pone en una gorountine
	fmt.Println("que lentooooo")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lentooooo(name string){
	letras := strings.Split(name,"") // Split divide los caracteres en un arreglo

	for _,letra := range letras{
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

}