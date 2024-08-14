package main

import (
	"fmt"
	"math"
)

type Quad struct {
	Q        Vec3
	U        Vec3
	V        Vec3
	W        Vec3
	Normal   Vec3
	D        float64
	Material Material
	AABB     AABB
}

func NewQuad(Q, U, V Vec3, Materal Material) Quad {
	n := U.Cross(V)
	Normal := n.Unit()
	D := Q.Dot(Normal)
	W := n.DivScalar(n.Dot(n))
	q := Quad{
		Q:        Q,
		U:        U,
		V:        V,
		W:        W,
		D:        D,
		Normal:   Normal,
		Material: Materal,
	}
	return q
}

func (q *Quad) SetAABB() {
	diag1 := NewAABBFromPoints(q.Q, q.Q.Add(q.U).Add(q.V))
	diag2 := NewAABBFromPoints(q.Q.Add(q.U), q.Q.Add(q.V))
	q.AABB = NewAABBFromAABB(diag1, diag2)
}

func (q Quad) Hit(r Ray, i Interval) (bool, HitInfo) {
	denominator := q.Normal.Dot(r.Direction)
	if math.Abs(denominator) < 0.0000001 {
		fmt.Println("denominator", denominator)
		return false, HitInfo{}
	}
	t := q.D - q.Normal.Dot(r.Direction)/denominator

	if i.Contains(t) {
		fmt.Println("t", t)
		return false, HitInfo{}
	}
	intersection := r.PointAt(t)
	planarHitpoint := intersection.Sub(q.Q)
	alpha, beta := q.W.Dot(planarHitpoint.Cross(q.V)), q.W.Dot(q.U.Cross(planarHitpoint))
	normal, frontFace := CalculateFaceNormal(r, q.Normal)

	fmt.Println(alpha, beta)
	if !IsInterior(alpha, beta) {
		return false, HitInfo{}
	}
	hitInfo := HitInfo{
		Point:     intersection,
		T:         t,
		Material:  q.Material,
		Normal:    normal,
		U:         alpha,
		V:         beta,
		FrontFace: frontFace,
	}
	return true, hitInfo
}
func (q Quad) BoundingBox() AABB {
	return q.AABB
}

func IsInterior(a, b float64) bool {
	unitInterval := Interval{Min: 0, Max: 1}
	if !unitInterval.Contains(a) || !unitInterval.Contains(b) {
		return false
	}
	return true
}
