package main

type Color struct {
	R, G, B int
}

func WriteColor(r, g, b float64) Color {
	r = max(0, min(1, r))
	g = max(0, min(1, g))
	b = max(0, min(1, b))
	ir := int(255.999 * r)
	ig := int(255.999 * g)
	ib := int(255.999 * b)
	return Color{ir, ig, ib}
}
