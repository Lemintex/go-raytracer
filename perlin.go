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

func (p Perlin) Noise(vec Vec3) float64 {
	u, v, w := vec.X-math.Floor(vec.X), vec.Y-math.Floor(vec.Y), vec.Z-math.Floor(vec.Z)
	u, v, w = u*u*(3-2*u), v*v*(3-2*v), w*w*(3-2*w)
	i, j, k := int(math.Floor(vec.X)), int(math.Floor(vec.Y)), int(math.Floor(vec.Z))
	c := make([][][]float64, 2)
	for di := 0; di < 2; di++ {
		c[di] = make([][]float64, 2)
		for dj := 0; dj < 2; dj++ {
			c[di][dj] = make([]float64, 2)
			for dk := 0; dk < 2; dk++ {
				c[di][dj][dk] = p.RandomFloats[p.PermX[(i+di)&255]^p.PermY[(j+dj)&255]^p.PermZ[(k+dk)&255]]
			}
		}
	}
	return PerlinInterp(c, u, v, w)
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

func PerlinInterp(c [][][]float64, u, v, w float64) float64 {
	accum := 0.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				accum += c[i][j][k] * (float64(i)*u + (1-float64(i))*(1-u)) * (float64(j)*v + (1-float64(j))*(1-v)) * (float64(k)*w + (1-float64(k))*(1-w))
			}
		}
	}
	return accum
}
