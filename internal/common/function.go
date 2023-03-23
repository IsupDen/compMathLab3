package common

import "math"

type function func(float64) float64

var (
	FunctionNames = []string{"sin(3x + 2)", "2x³ - 9x² - 7x + 11", "cos(2x)² - 3sin(x)", "|cos(x) - x³| - 4x", "(x² + 1) * (5x - 1)", "2 / √(1 - x²)", "1 / ln³x"}
	functions     = map[string]function{
		"sin(3x + 2)": func(x float64) float64 {
			return math.Sin(3*x + 2)
		},
		"2x³ - 9x² - 7x + 11": func(x float64) float64 {
			return 2*math.Pow(x, 3) - 9*math.Pow(x, 2) - 7*x + 11
		},
		"cos(2x)² - 3sin(x)": func(x float64) float64 {
			return math.Pow(math.Cos(2*x), 2) - 3*math.Sin(x)
		},
		"|cos(x) - x³| - 4x": func(x float64) float64 {
			return math.Abs(math.Cos(x)-math.Pow(x, 3)) - 4*x
		},
		"(x² + 1) * (5x - 1)": func(x float64) float64 {
			return (math.Pow(x, 2) + 1) * (5*x - 1)
		},
		"2 / √(1 - x²)": func(x float64) float64 {
			return 2 / math.Sqrt(1-math.Pow(x, 2))
		},
		"1 / ln³x": func(x float64) float64 {
			return 1 / math.Pow(math.Log(x), 3)
		},
	}
	primitives = map[string]function{
		"2 / √(1 - x²)": func(x float64) float64 {
			return 2 * math.Asin(x)
		},
		"1 / ln³x": func(x float64) float64 {
			return 1 / 2 / math.Pow(math.Log(x), 2)
		},
	}
)
