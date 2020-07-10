package main 

import(
	"fmt"
	"time"
)

func imprimirPing(c chan string) {
	var cantador int
	for{
		//recibiendo valores atraves del canal
		cantador++
		fmt.Println(<-c, " ", cantador)
		time.Sleep(time.Second * 1)
	}
}

func enviarPing(c chan string) {
	for{
		//enviamos valores atraves del canal
		c <- "Ping"
	}
}

func main() {
	
	//creamos un canal
	ch := make(chan string)

	//llamamos las fuciones como gorountine
	go enviarPing(ch)
	go imprimirPing(ch)

	//esperamos a que lea algo para finalizar el main y parar las gorountine
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin...")
}

/*
	aunque ambas gorutinas estan en for infinitas, automaticamente se
	paran para recibir y enviar informacion atraves de los canales. 

	Prime el ciclo infinito se para para enviar la informacion
	Luego el otro ciclo infinita recibe la informacin y luego se vuelve a
	para para recibir otra vez

	El programa se para cuando se escandea algo para cerrar el main
*/