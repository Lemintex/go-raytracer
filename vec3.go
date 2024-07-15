package main

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
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
