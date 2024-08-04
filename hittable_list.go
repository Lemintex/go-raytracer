package main

type HittableList struct {
	Objects []Hittable
	AABB    AABB
}

func (hl *HittableList) Add(h Hittable) {
	hl.Objects = append(hl.Objects, h)
	hl.AABB = NewAABBFromAABB(hl.AABB, h.BoundingBox())
}

func (hl *HittableList) Clear() {
	hl.Objects = nil
}

func (hl HittableList) Hit(r Ray, i Interval) (bool, HitInfo) {
	hitAnything := false
	closestSoFar := i.Max
	hitInfo := HitInfo{}
	for _, object := range hl.Objects {
		hit, tempHitInfo := object.Hit(r, Interval{i.Min, closestSoFar})
		if hit {
			hitAnything = true
			closestSoFar = tempHitInfo.T
			hitInfo = tempHitInfo
		}
	}
	return hitAnything, hitInfo
}
