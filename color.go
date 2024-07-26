package main

type Color struct {
	R, G, B int
}

func WriteColor(r, g, b float64) Color {
	r = max(0, min(1, r))
	g = max(0, min(1, g))
	b = max(0, min(1, b))

	intensity := Interval{0.0, 0.999}
	ir := int(256 * intensity.Clamp(r))
	ig := int(256 * intensity.Clamp(g))
	ib := int(256 * intensity.Clamp(b))
	return Color{ir, ig, ib}
}
