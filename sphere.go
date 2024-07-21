package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, HitInfo) {
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
	if root < tMin || tMax < root {
		root = (-h + sqrtDiscriminant) / a
		if root < tMin || tMax < root {
			return false, HitInfo{}
		}
	}

	point := r.PointAt(root)
	normal := point.Sub(s.Center).DivScalar(s.Radius)
	hitInfo := HitInfo{
		Point:  point,
		Normal: normal,
		T:      root,
	}
	return true, hitInfo
}

func (s Sphere) CalculateFaceNormal(r Ray, outwardNormal Vec3) (Vec3, bool) {
	frontFace := r.Direction.Dot(outwardNormal) < 0
	return outwardNormal, frontFace
}
