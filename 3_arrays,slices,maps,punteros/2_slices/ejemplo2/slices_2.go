package main

import "fmt"

func main() {
	var min int
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for i, v := range x {
		if i == 0 || v < min {
			min = v
		}

	}

	fmt.Println(min)
}
