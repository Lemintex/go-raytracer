package main

type HittableList struct {
	Objects []Hittable
}

func (hl *HittableList) Add(h Hittable) {
	hl.Objects = append(hl.Objects, h)
}

func (hl *HittableList) Clear() {
	hl.Objects = nil
}

func (hl HittableList) Hit(r Ray, tMin float64, tMax float64) (bool, HitInfo) {
	hitAnything := false
	closestSoFar := tMax
	hitInfo := HitInfo{}
	for _, object := range hl.Objects {
		hit, tempHitInfo := object.Hit(r, tMin, closestSoFar)
		if hit {
			hitAnything = true
			closestSoFar = tempHitInfo.T
			hitInfo = tempHitInfo
		}
	}
	return hitAnything, hitInfo
}
