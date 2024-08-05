package main

type BvhNode struct {
	Left, Right Hittable
	AABB        AABB
}

func (node BvhNode) Hit(r Ray, i Interval) bool {
	if !node.AABB.Hit(r, i) {
		return false
	}
	var hitLeft, hitRight bool
	if node.Left != nil {
		hitLeft, _ = node.Left.Hit(r, i)
	}

	if node.Right != nil {
		hitRight, _ = node.Right.Hit(r, i)
	}

	return hitLeft || hitRight
}
