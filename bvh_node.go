package main

import (
	"sort"
)

type BvhNode struct {
	Left, Right Hittable
	AABB        AABB
}

func (n BvhNode) NewBvhNode(aaab []Hittable, start, end int) BvhNode {
	bbox := AABB{}
	for i := start; i < end; i++ {
		bbox = NewAABBFromAABB(bbox, aaab[i].BoundingBox())
	}
	axis := bbox.LongestAxis()
	var comparator func(a, b Hittable) bool
	if axis == 0 {
		comparator = n.BoxCompareX
	} else if axis == 1 {
		comparator = n.BoxCompareY
	} else {
		comparator = n.BoxCompareZ
	}
	objectSpan := end - start

	if objectSpan == 1 {
		n.Left = aaab[start]
		n.Right = aaab[start]
	} else if objectSpan == 2 {
		n.Left = aaab[start]
		n.Right = aaab[start+1]
	} else {
		sort.Slice(aaab[start:end], func(i, j int) bool {
			return comparator(aaab[i], aaab[j])
		})
		mid := start + objectSpan/2
		n.Left = n.NewBvhNode(aaab, start, mid)
		n.Right = n.NewBvhNode(aaab, mid, end)
	}
	n.AABB = NewAABBFromAABB(n.Left.BoundingBox(), n.Right.BoundingBox())
	return n
}

func (node BvhNode) Hit(r Ray, i Interval) (bool, HitInfo) {
	if !node.AABB.Hit(r, i) {
		return false, HitInfo{}
	}
	var hitLeft, hitRight bool
	hitInfo := HitInfo{}
	if node.Left != nil {
		hitLeft, hitInfo = node.Left.Hit(r, i)
	}

	if node.Right != nil {
		var temp float64
		var tempHitInfo HitInfo
		if hitLeft {
			temp = hitInfo.T
		} else {
			temp = i.Max
		}
		hitRight, tempHitInfo = node.Right.Hit(r, Interval{i.Min, temp})
		if hitRight {
			hitInfo = tempHitInfo
		}
	}

	return hitLeft || hitRight, hitInfo
}

func (n BvhNode) BoxCompare(a, b Hittable, axisIndex int) bool {
	aAxisInterval := a.BoundingBox().AxisInterval(axisIndex)
	bAxisInterval := b.BoundingBox().AxisInterval(axisIndex)
	return aAxisInterval.Min < bAxisInterval.Min

}
func (n BvhNode) BoxCompareX(a, b Hittable) bool {
	boxA := a.BoundingBox()
	boxB := b.BoundingBox()
	return boxA.X.Min < boxB.X.Min
}

func (n BvhNode) BoxCompareY(a, b Hittable) bool {
	boxA := a.BoundingBox()
	boxB := b.BoundingBox()
	return boxA.Y.Min < boxB.Y.Min
}

func (n BvhNode) BoxCompareZ(a, b Hittable) bool {
	boxA := a.BoundingBox()
	boxB := b.BoundingBox()
	return boxA.Z.Min < boxB.Z.Min
}

func (n BvhNode) BoundingBox() AABB {
	return n.AABB
}
