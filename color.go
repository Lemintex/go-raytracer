package main

import "math"

type Color struct {
	R, G, B int
}

func linearToGamma(x float64) float64 {
	if x > 0 {
		return math.Sqrt(x)
	}
	return 0
}

func WriteColor(r, g, b float64) Color {
	r = max(0, min(1, r))
	g = max(0, min(1, g))
	b = max(0, min(1, b))

	r = linearToGamma(r)
	g = linearToGamma(g)
	b = linearToGamma(b)

	intensity := Interval{0.0, 0.999}
	ir := int(256 * intensity.Clamp(r))
	ig := int(256 * intensity.Clamp(g))
	ib := int(256 * intensity.Clamp(b))
	return Color{ir, ig, ib}
}
