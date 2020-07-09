package main 

// de esta forma lee linea por linea en ves de todo de golpe

import(
	"fmt"
	"bufio"
	"os"
)

func leer() bool{
	archivo, err := os.Open("./hola") // abre el archivo

	defer func() {
		archivo.Close()
		fmt.Println("defer")

		r := recover()
		fmt.Println(r)	
	}()

	if err != nil{
		panic(err)
	}

	lector := bufio.NewScanner(archivo) // aqui es donde lo va a leer

	var i int
	for lector.Scan(){

		i++
		linea := lector.Text()

		fmt.Printf("%d: %v\n", i, linea )
	}

	return true

	fmt.Println("nunca me ejecuto")

	return true
}

func ejecutar() {
	arc := leer()
	fmt.Println(arc)
	
}

func main() {

	ejecutar()
}	

/*
	defer 
		se ejecuta siempre al terminal la funcion por eso es 
		bueno para cerrar los archivos por seguridad

	panic
		se usa para imprimir un error e identicifar en que linea
		esta pasando

	recover
		detiene el panic, la ejecucion continua eceptuando el panic
		y tambien puede imprimir el tipo de error

*/