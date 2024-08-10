package main

import (
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
