package main

type Color struct {
	R, G, B int
}

func WriteColor(r, g, b float32) Color {
	ir := int(255.999 * r)
	ig := int(255.999 * g)
	ib := int(255.999 * b)
	return Color{ir, ig, ib}
}
