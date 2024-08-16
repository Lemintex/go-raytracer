package main

type Interval struct {
	Min, Max float64
}

func NewInterval(a, b Interval) (float64, float64) {
	var min, max float64
	if a.Min < b.Min {
		min = a.Min
	} else {
		min = b.Min
	}
	if a.Max > b.Max {
		max = a.Max
	} else {
		max = b.Max
	}
	return min, max
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

func (i *Interval) Expand(delta float64) {
	padding := delta / 2
	i.Min -= padding
	i.Max += padding
}
