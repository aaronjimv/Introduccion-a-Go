package main 

import(
	"fmt"
)

func main() {
	
	fmt.Println("Hola mundo")

	edad := 21
	fmt.Printf("mi edad es %d\n", edad) // los % son verbos, debo investigar eso

	nombre := "Aaron"
	fmt.Printf("mi nombre es: %v y tengo %v\n" , nombre,edad)

	estado := false
	fmt.Printf("el estado es %t\n" , estado)

	precio := 20.5
	fmt.Printf("el precio es de %f $\n" , precio)

	fmt.Printf("mi nombre es: %s\n" , "Aaron")

	const nom, num = "Aaron",21 // al parecer con const se puede declarar diferentes tipos a la vez
	fmt.Print(nom, " is ", num, " years old.\n")
	
	//leer de teclado

	var nombre2 string
	var edad2 int
	fmt.Println("hola, coloca tu nombre: ")
	fmt.Scanf("%s\n" , &nombre2) // scanf es para leer, %s es para string y &nombre2 es donde se va a guardar
	fmt.Println("coloca tu edad: ")
	fmt.Scanf("%d\n" , &edad2) // %d para int
	fmt.Printf("te llamas %s y tines %d\n" , nombre2,edad2)
}