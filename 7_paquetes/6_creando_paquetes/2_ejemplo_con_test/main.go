package main

import (
	"fmt"

	m "main/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}

	avg := m.Average(xs)
	fmt.Println(avg)

	min := m.Min(xs)
	fmt.Println(min)

	max := m.Max(xs)
	fmt.Println(max)
}
