package main

type Quad struct {
	Q        Vec3
	U        Vec3
	V        Vec3
	Material Material
	AABB     AABB
}

func NewQuad(Q, U, V Vec3, Materal Material) Quad {
	q := Quad{
		Q:        Q,
		U:        U,
		V:        V,
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
	return false, HitInfo{}
}
func (q Quad) BoundingBox() AABB {
	return q.AABB
}
