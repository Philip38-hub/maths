package mathskill

func Average(x []float64) float64 {
	sum := 0.0
	counter := 0.0
	for _, num := range x {
		sum += num
		counter++
	}
	return sum / counter
}