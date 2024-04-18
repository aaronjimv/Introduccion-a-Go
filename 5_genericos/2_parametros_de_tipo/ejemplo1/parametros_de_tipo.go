package main

import (
	"fmt"
)

type MyInt int
type MyIntV2 int

func main() {
	var num1 MyInt = 2
	var num2 MyInt = 6

	var num3 MyIntV2 = 3
	var num4 MyIntV2 = 4

	fmt.Println(sum(2, 4, 67))
	fmt.Println(sum(2.0, 4.88, 67.1))
	fmt.Println(sum(25, 3.89, 67.1, 8))
	fmt.Println(sum(num1, num2))
	fmt.Println(sum(num3, num4))
	// fmt.Println(sum(3, 5.78, num1, num3)) error
}

func sum[T ~int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}
