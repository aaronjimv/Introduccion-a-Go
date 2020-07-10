package main 

import(
	"fmt"
	"./ejemplo" // ./ -> porque el paquete esta en la misma carpeta
)

func main() {
	fmt.Println(ejemplo.Hola)
	fmt.Println(ejemplo.HolaMundo())
}