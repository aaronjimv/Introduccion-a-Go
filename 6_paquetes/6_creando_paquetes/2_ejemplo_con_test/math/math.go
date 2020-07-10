package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// returns the maximum number in a slice of float64s.
func Max(xs []float64) float64 {
	var max float64
	for i, v := range xs {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

// returns the minimum number in a slice of float64s.
func Min(xs []float64) float64 {
	var min float64
	for i, v := range xs {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}