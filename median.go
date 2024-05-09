package mathskill

func Median(x []float64) float64 {
	data := Sort(x)
	out := 0.0
	if len(data)%2 == 1 {
		out = data[(len(x) / 2)]
	} else {
		out = (data[len(data)/2] + data[(len(data)/2)-1]) / 2
	}
	return out
}