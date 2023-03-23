package common

import m "lab3/internal/methods"

type method func(m.Integral, int) float64

var (
	MethodNames = []string{"Rectangle (left)", "Rectangle (middle)", "Rectangle (right)", "Trapezoid", "Simpson"}
	methods     = map[string]method{
		"Rectangle (left)":   m.Rectangle(1),
		"Rectangle (middle)": m.Rectangle(2),
		"Rectangle (right)":  m.Rectangle(3),
		"Trapezoid":          m.Trapezoid,
		"Simpson":            m.Simpson,
	}
)
