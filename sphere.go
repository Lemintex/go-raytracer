package main

import "math"

type Sphere struct {
	Center   Vec3
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r Ray, i Interval) (bool, HitInfo) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	h := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c
	if discriminant < 0 {
		return false, HitInfo{}
	}

	sqrtDiscriminant := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-h - sqrtDiscriminant) / a
	if !i.Contains(root) {
		root = (-h + sqrtDiscriminant) / a
		if !i.Contains(root) {
			return false, HitInfo{}
		}
	}

	point := r.PointAt(root)
	normal := point.Sub(s.Center).DivScalar(s.Radius)
	normal, frontFace := s.CalculateFaceNormal(r, normal)
	hitInfo := HitInfo{
		Point:     point,
		Normal:    normal,
		Material:  s.Material, //FIXME: Material is not defined
		T:         root,
		FrontFace: frontFace, //FIXME: FrontFace is not defined
	}
	return true, hitInfo
}

func (s Sphere) CalculateFaceNormal(r Ray, outwardNormal Vec3) (Vec3, bool) {
	frontFace := r.Direction.Dot(outwardNormal) < 0
	if !frontFace {
		outwardNormal = outwardNormal.Neg()
	}
	return outwardNormal, frontFace
}
