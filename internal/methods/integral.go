package methods

import "math"

type Integral interface {
	Function(float64) float64
	Primitive(float64) float64
	A() float64
	B() float64
}

func value(i Integral, x float64) float64 {
	val := i.Function(x)
	var sign int
	if math.Signbit(val) {
		sign = -1
	} else {
		sign = 1
	}
	if math.IsNaN(val) || math.IsInf(val, sign) {
		val = i.Primitive(x)
		if math.Signbit(val) {
			sign = -1
		} else {
			sign = 1
		}
		if math.IsNaN(val) || math.IsInf(val, sign) {
			return math.NaN()
		}
		return 0
	}
	return val
}
