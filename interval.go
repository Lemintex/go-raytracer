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

func (i Interval) Clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	}
	if x > i.Max {
		return i.Max
	}
	return x
}
func (i Interval) Expand(x float64) Interval {
	return Interval{i.Min - x, i.Max + x}
}
