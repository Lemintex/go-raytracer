package main

type HitInfo struct {
	Point  Vec3
	Normal Vec3
	T      float64
}

type Hittable interface {
	Hit(r Ray, i Interval) (bool, HitInfo)
}
