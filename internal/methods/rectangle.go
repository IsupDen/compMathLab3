package methods

func getX(number int) func(float64, float64) float64 {
	switch number {
	case 1:
		return func(a float64, b float64) float64 {
			return a
		}
	case 2:
		return func(a float64, b float64) float64 {
			return (a + b) / 2
		}
	case 3:
		return func(a float64, b float64) float64 {
			return b
		}
	}
	return nil
}

func Rectangle(number int) func(Integral, int) float64 {
	return func(integral Integral, n int) float64 {
		X := getX(number)
		h := (integral.B() - integral.A()) / float64(n)
		sum := 0.0
		for i := integral.A(); i < integral.B(); i += h {
			sum += value(integral, X(i, i+h))
		}
		return sum * h
	}
}
