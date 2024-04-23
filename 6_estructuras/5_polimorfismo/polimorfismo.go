package main

import "fmt"

// Define una interfaz para métodos de pago
type PayMethod interface {
	// El método Pay es un comportamiento común para todos los métodos de pago
	Pay()
}

// Define una estructura Paypal que implementa la interfaz PayMethod
type Paypal struct{}

// Implementa el método Pay para Paypal
func (p Paypal) Pay() {
	fmt.Println("Pagado por paypal")
}

// Define una estructura Cash que implementa la interfaz PayMethod
type Cash struct{}

// Implementa el método Pay para Cash
func (c Cash) Pay() {
	fmt.Println("Pagado por efectivo")
}

// Define una estructura CreditCard que implementa la interfaz PayMethod
type CreditCard struct{}

// Implementa el método Pay para CreditCard
func (c CreditCard) Pay() {
	fmt.Println("Pagado por tarjeta de crédito")
}

// Una función de fábrica que devuelve una instancia de PayMethod según el método de pago ingresado
func Factory(method uint) PayMethod {
	switch method {
	case 1:
		// Devuelve una instancia de Paypal
		return Paypal{}
	case 2:
		// Devuelve una instancia de Cash
		return Cash{}
	case 3:
		// Devuelve una instancia de CreditCard
		return CreditCard{}
	default:
		// Devuelve nil si el método de pago ingresado es inválido
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite un número de método de pago")
	fmt.Println("\t 1:Paypal 2:Efectivo 3:Tarjeta de crédito")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("debe digitar un número válido")
	}
	if method > 3 {
		panic("debe digitar un número válido")
	}

	// Crea una instancia de PayMethod utilizando la función de fábrica
	payMethod := Factory(method)

	// Llama al método Pay en la instancia de PayMethod, lo que ejecutará el método de pago correcto
	payMethod.Pay()
}

/*
Este código demuestra el polimorfismo al
definir una interfaz PayMethod con un solo
método Pay().

Tres estructuras, Paypal, Cash y CreditCard,
implementan esta interfaz proporcionando su propia
implementación del método Pay().

La función Factory devuelve una instancia
de una de estas estructuras según el método
de pago ingresado, y la función main utiliza
esta instancia para llamar al método Pay(),
lo que ejecutará el método de pago correcto.
*/
