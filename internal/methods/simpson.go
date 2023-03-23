package methods

func Simpson(integral Integral, n int) float64 {
	h := (integral.B() - integral.A()) / float64(n)
	sum := 0.0
	for i := integral.A() + h; i < integral.B(); i += 2 * h {
		sum += 4 * value(integral, i)
	}
	for i := integral.A() + 2*h; i < integral.B(); i += 2 * h {
		sum += 2 * value(integral, i)
	}
	return h / 3 * (value(integral, integral.A()) + value(integral, integral.B()) + sum)
}
