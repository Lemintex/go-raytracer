package main

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3) AddScalar(s float64) Vec3 {
	return Vec3{v.X + s, v.Y + s, v.Z + s}
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vec3) Div(v2 Vec3) Vec3 {
	return Vec3{v.X / v2.X, v.Y / v2.Y, v.Z / v2.Z}
}

func (v Vec3) MulScalar(s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3) DivScalar(s float64) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{v.Y*v2.Z - v.Z*v2.Y, v.Z*v2.X - v.X*v2.Z, v.X*v2.Y - v.Y*v2.X}
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

func (v Vec3) Unit() Vec3 {
	return v.DivScalar(v.Length())
}

func SampleSquare() Vec3 {
	return Vec3{RandomFloatBetweenMinAndMax(-0.5, 0.5), RandomFloatBetweenMinAndMax(-0.5, 0.5), 0}
}

func RandomVec3() Vec3 {
	return Vec3{RandomFloat(), RandomFloat(), RandomFloat()}
}

func RandomVec3BetweenMinAndMax(min, max float64) Vec3 {
	return Vec3{RandomFloatBetweenMinAndMax(min, max), RandomFloatBetweenMinAndMax(min, max), RandomFloatBetweenMinAndMax(min, max)}
}

func RandomVec3InUnitSphere() Vec3 {
	for range 100 {
		p := RandomVec3BetweenMinAndMax(-1, 1)
		if p.LengthSquared() < 1 {
			return p
		}
	}
	return Vec3{}
}

func RandomUnitVec3() Vec3 {
	return RandomVec3InUnitSphere().Unit()
}

func RandomVec3OnHemisphere(normal Vec3) Vec3 {
	inUnitSphere := RandomVec3InUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Neg()
}

func (v Vec3) IsNearZero() bool {
	s := 1e-6
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.MulScalar(2 * v.Dot(n)))
}

func (v Vec3) Refract(n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(v.Neg().Dot(n), 1.0)
	rOutPerp := v.Add(n.MulScalar(cosTheta)).MulScalar(etaiOverEtat)
	rOutParallel := n.MulScalar(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}
