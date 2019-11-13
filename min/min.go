package min

func Min(xs []float64) float64 {
	y := xs[0]
	for i, x := range xs {
		if x < y {
			y = xs[i]
		}
	}
	return y
}
