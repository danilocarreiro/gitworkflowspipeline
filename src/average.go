package src

// Average between numbers passed
func Average(numbers ...float64) float64 {
	var total float64 = 0.0

	for _, number := range numbers {
		total += number
	}

	return (total / float64(len(numbers)))
}
