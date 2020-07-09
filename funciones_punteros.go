package main 

import "fmt"

/*
	When we call a function that takes an argument, that argument is copied to the func‐
	tion:
		func zero(x int) {
			x = 0
		}
		func main() {
			x := 5
			zero(x)
			fmt.Println(x) // x is still 5
		}

	In this program, the zero function will not modify the original x variable in the main
	function. But what if we wanted to? One way to do this is to use a special data type
	known as a pointer:
*/
func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

/*
	Pointers reference a location in memory where a value is stored 
	rather than the value itself. By using a pointer (*int), the zero 
	function is able to modify the original variable.
*/
