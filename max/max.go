package max

func Max(xs []float64) float64 {
	y := xs[0]
	for i, x := range xs {
		if y < xs[i] {
			y = x
		}
	}
	return y
}
