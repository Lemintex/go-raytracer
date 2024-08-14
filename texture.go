package main

import (
	"image"
	"math"
)

type Texture interface {
	Value(u, v float64, p Vec3) Vec3
}

type SolidColor struct {
	Albedo Vec3
}

func (c SolidColor) Value(u, v float64, p Vec3) Vec3 {
	return c.Albedo
}

type CheckerTexture struct {
	Scale     float64
	Even, Odd Texture
}

func (c CheckerTexture) Value(u, v float64, p Vec3) Vec3 {
	invScale := 1 / c.Scale
	x := int(math.Floor(invScale * p.X))
	y := int(math.Floor(invScale * p.Y))
	z := int(math.Floor(invScale * p.Z))
	isEven := (x+y+z)%2 == 0
	if isEven {
		return c.Even.Value(u, v, p)
	} else {
		return c.Odd.Value(u, v, p)
	}
}

type ImageTexture struct {
	img image.Image
}

func (i ImageTexture) Value(u, v float64, p Vec3) Vec3 {
	if i.img == nil {
		return Vec3{1, 0, 1}
	}
	bounds := i.img.Bounds()
	x := int(u * float64(bounds.Max.X))
	y := int((1 - v) * float64(bounds.Max.Y))
	r, g, b, _ := i.img.At(x, y).RGBA()
	return Vec3{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535}
}

type NoiseTexture struct {
	Noise Perlin
	Scale float64
}

func (n NoiseTexture) Value(u, v float64, p Vec3) Vec3 {
	return Vec3{1, 1, 1}.MulScalar(n.Noise.Turb(p, 7))
}
