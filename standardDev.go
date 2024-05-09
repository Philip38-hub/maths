package mathskill

import "math"

func StandardDev(x []float64) float64 {
	out := Variance(x)
	return math.Sqrt(out)
}
