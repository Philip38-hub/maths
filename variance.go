package mathskill

func Variance(x []float64) float64 {
	mean := Average(x)
	out := []float64{}

	for _, num := range x {
		dev := num - mean
		out = append(out, dev*dev)
	}
	return Average(out)
}
