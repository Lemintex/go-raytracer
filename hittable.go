package main

type HitInfo struct {
	Point     Vec3
	Normal    Vec3
	Material  Material
	T         float64
	FrontFace bool
}

type Hittable interface {
	Hit(r Ray, i Interval) (bool, HitInfo)
	BoundingBox() AABB
}
