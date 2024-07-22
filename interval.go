package main

type Interval struct {
	Min, Max float64
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) Contains(x float64) bool {
	return x >= i.Min && x <= i.Max
}

func (i Interval) Surrounds(x float64) bool {
	return x > i.Min && x < i.Max
}
