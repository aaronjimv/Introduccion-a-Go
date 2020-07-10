package main 

import "fmt"

func main() {
	
	var x,y int
	fmt.Println("ingrese un numero: ")
	fmt.Scanf("%d\n" , &x)
	fmt.Println("ingrese otro numero: ")
	fmt.Scanf("%d\n" , &y)

	if x < y{
		fmt.Printf("%d es mayor que %d\n", y,x)
	}else if x > y{
		fmt.Printf("%d es mayor que %d\n", x,y)
	}else{
		fmt.Printf("%d es igual que %d\n", x,y)
	}
}