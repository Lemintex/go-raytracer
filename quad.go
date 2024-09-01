package main

import (
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
	D := Normal.Dot(Q)
	W := n.DivScalar(n.Length())
	q := Quad{
		Q:        Q,
		U:        U,
		V:        V,
		W:        W,
		Normal:   Normal,
		D:        D,
		Material: Materal,
	}
	q.SetAABB()
	return q
}

func (q *Quad) SetAABB() {
	d1 := NewAABBFromPoints(q.Q, q.Q.Add(q.U).Add(q.V))
	d2 := NewAABBFromPoints(q.Q.Add(q.U), q.Q.Add(q.V))
	q.AABB = NewAABBFromAABB(d1, d2)
}

func (q Quad) Hit(r Ray, i Interval) (bool, HitInfo) {
	denom := q.Normal.Dot(r.Direction)
	if math.Abs(denom) < 0.0001 {
		return false, HitInfo{}
	}
	t := (q.D - q.Normal.Dot(r.Origin)) / denom
	if !i.Contains(t) {
		return false, HitInfo{}
	}
	p := r.PointAt(t)
	du := p.Sub(q.Q)
	u := du.Dot(q.U)
	if u < 0 || u > q.U.Length() {
		return false, HitInfo{}
	}
	dv := p.Sub(q.Q)
	v := dv.Dot(q.V)
	if v < 0 || v > q.V.Length() {
		return false, HitInfo{}
	}
	return true, HitInfo{
		T:        t,
		Point:    p,
		Normal:   q.Normal,
		Material: q.Material,
	}
}
func (q Quad) BoundingBox() AABB {
	return q.AABB
}

func IsInterior(a, b float64) bool {
	return a > 0 && b > 0 || a < 0 && b < 0
}
