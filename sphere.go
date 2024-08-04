package main

import "math"

type Sphere struct {
	IsMoving    bool
	CenterInit  Vec3
	CenterFinal Vec3
	Radius      float64
	Material    Material
	AABB        AABB
}

func NewStationarySphere(center Vec3, radius float64, material Material) Sphere {
	radiusVec := Vec3{radius, radius, radius}
	aabb := NewAABBFromPoints(center.Sub(radiusVec), center.Add(radiusVec))
	return Sphere{
		IsMoving:    false,
		CenterInit:  center,
		CenterFinal: center,
		Radius:      radius,
		Material:    material,
		AABB:        aabb,
	}
}

func NewMovingSphere(centerInit, centerFinal Vec3, radius float64, material Material) Sphere {
	radiusVec := Vec3{radius, radius, radius}
	box1 := NewAABBFromPoints(centerInit.Sub(radiusVec), centerInit.Add(radiusVec))
	box2 := NewAABBFromPoints(centerFinal.Sub(radiusVec), centerFinal.Add(radiusVec))
	bbox := NewAABBFromAABB(box1, box2)
	return Sphere{
		IsMoving:    true,
		CenterInit:  centerInit,
		CenterFinal: centerFinal,
		Radius:      radius,
		Material:    material,
		AABB:        bbox,
	}
}

func (s Sphere) Hit(r Ray, i Interval) (bool, HitInfo) {
	oc := r.Origin.Sub(s.Center(r.Time))
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
	normal := point.Sub(s.Center(r.Time)).DivScalar(s.Radius)
	normal, frontFace := s.CalculateFaceNormal(r, normal)
	hitInfo := HitInfo{
		Point:     point,
		Normal:    normal,
		Material:  s.Material,
		T:         root,
		FrontFace: frontFace,
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

func (s Sphere) Center(time float64) Vec3 {
	if !s.IsMoving {
		return s.CenterInit
	}
	return s.CenterInit.Add(s.CenterFinal.Sub(s.CenterInit).MulScalar(time))
}

func (s Sphere) BoundingBox() AABB {
	return s.AABB
}
