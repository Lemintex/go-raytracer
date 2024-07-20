package main

import "math"

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) PointAt(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

func (r Ray) Color() Vec3 {
	t := r.HitSphere(Vec3{0, 0, -1}, 0.5)
	if t > 0 {
		normal := r.PointAt(t).Sub(Vec3{0, 0, -1}).Unit()
		return normal.Add(Vec3{1, 1, 1}).MulScalar(0.5)
	}
	unitDir := r.Direction.Unit()
	a := 0.5 * (unitDir.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(a))
}

func (r Ray) HitSphere(center Vec3, radius float64) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.LengthSquared()
	h := oc.Dot(r.Direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := h*h - a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-h - math.Sqrt(discriminant)) / a
}
