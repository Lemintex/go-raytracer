package main

import "math"

type Perlin struct {
	PointCount   int
	RandomFloats []float64
	PermX        []int
	PermY        []int
	PermZ        []int
}

func NewPerlin() Perlin {
	p := Perlin{}
	p.PointCount = 256
	p.RandomFloats = make([]float64, p.PointCount)
	for i := range p.PointCount {
		p.RandomFloats[i] = RandomFloat()
	}
	p.PermX = PerlinGeneratePerm()
	p.PermY = PerlinGeneratePerm()
	p.PermZ = PerlinGeneratePerm()
	return p
}

func (p Perlin) Noise(v Vec3) float64 {
	i := int(4*v.X) & 255
	j := int(4*v.Y) & 255
	k := int(4*v.Z) & 255
	return p.RandomFloats[p.PermX[i]^p.PermY[j]^p.PermZ[k]]
}

func (p Perlin) Turb(v Vec3, depth int) float64 {
	accum := 0.0
	tempV := v
	weight := 1.0
	for i := 0; i < depth; i++ {
		accum += weight * p.Noise(tempV)
		weight *= 0.5
		tempV = tempV.MulScalar(2)
	}
	return math.Abs(accum)
}

func PerlinGeneratePerm() []int {
	p := make([]int, 256)
	for i := range 256 {
		p[i] = i
	}
	p = PerlinPermute(p, 256)
	return p
}

func PerlinPermute(p []int, n int) []int {
	for i := n - 1; i > 0; i-- {
		target := int(float64(RandomIntBetweenMinAndMax(0, i)))
		p[i], p[target] = p[target], p[i]
	}
	return p
}
