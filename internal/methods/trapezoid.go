package methods

func Trapezoid(integral Integral, n int) float64 {
	h := (integral.B() - integral.A()) / float64(n)
	sum := 0.0
	for i := integral.A() + h; i < integral.B(); i += h {
		sum += value(integral, i)
	}
	return h / 2 * (value(integral, integral.A()) + value(integral, integral.B()) + 2*sum)
}
