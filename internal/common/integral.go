package common

import (
	"fmt"
	"math"
)

var (
	Quit    chan bool
	Results = make(chan *Result)
)

type Integral struct {
	function  function
	method    method
	primitive function
	a         float64
	b         float64
	accuracy  float64
}

func (i *Integral) Solve() {
	Quit = make(chan bool)
	go func() {
		n := 4
		i0 := i.method(i, n)
		n *= 2
		i1 := i.method(i, n)
		if math.IsNaN(i1) || math.IsNaN(i0) {
			Results <- &Result{
				IntegralValue: math.NaN(),
				N:             0,
			}
			return
		}
		for math.Abs(i0-i1) >= i.accuracy {
			fmt.Println(i0)
			select {
			case <-Quit:
				Results <- nil
				return
			default:
				n *= 2
				i0 = i1
				i1 = i.method(i, n)
				if math.IsNaN(i1) {
					Results <- &Result{
						IntegralValue: math.NaN(),
						N:             0,
					}
					return
				}

			}
		}
		Results <- &Result{
			IntegralValue: i1,
			N:             n,
		}
	}()
}

type Result struct {
	IntegralValue float64
	N             int
}

func (i *Integral) Function(x float64) float64 {
	return i.function(x)
}

func (i *Integral) Primitive(x float64) float64 {
	return i.primitive(x)
}

func (i *Integral) A() float64 {
	return i.a
}

func (i *Integral) B() float64 {
	return i.b
}

func (i *Integral) Accuracy() float64 {
	return i.accuracy
}

func (i *Integral) GetFunction() function {
	return i.function
}

func (i *Integral) GetMethod() method {
	return i.method
}

func (i *Integral) SetFunction(functionName string) {
	i.function = functions[functionName]
}

func (i *Integral) SetA(a float64) {
	i.a = a
}

func (i *Integral) SetB(b float64) {
	i.b = b
}

func (i *Integral) SetMethod(methodName string) {
	i.method = methods[methodName]
}

func (i *Integral) SetAccuracy(accuracy float64) {
	i.accuracy = accuracy
}

func (i *Integral) GetPrimitive() function {
	return i.primitive
}

func (i *Integral) SetPrimitive(functionName string) {
	i.primitive = primitives[functionName]
}
