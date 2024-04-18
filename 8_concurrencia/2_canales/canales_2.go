package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4,0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
	fmt.Println(y, x, x+y)
}

/*
	las go runtine operan como pilas al llegar a la linea 17 se ejecuta en el fondo
	y salta a la 18 luego esta se ejecuta en el fondo

	el primer el primer mensaje lo recibira del ultimo canal

	Last In Firth Out
*/
