package main

type HittableList struct {
	Objects []Hittable
	AABB    AABB
	BvhRoot BvhNode
}

func (hl *HittableList) Add(h Hittable) {
	hl.Objects = append(hl.Objects, h)
	hl.AABB = NewAABBFromAABB(hl.AABB, h.BoundingBox())
}

func (hl *HittableList) Clear() {
	hl.Objects = nil
}

func (hl HittableList) Hit(r Ray, i Interval) (bool, HitInfo) {
	return hl.BvhRoot.Hit(r, i)
}

func (hl *HittableList) BuildBVH() {
	hl.BvhRoot = hl.BvhRoot.NewBvhNode(hl.Objects, 0, len(hl.Objects))
}
