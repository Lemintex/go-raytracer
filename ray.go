package main

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) PointAt(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

func (r Ray) Color(world HittableList) Vec3 {
	didHit, hit := world.Hit(r, 0, 100)
	if didHit {
		normal := r.PointAt(hit.T).Sub(Vec3{0, 0, -1}).Unit()
		return normal.Add(Vec3{1, 1, 1}).MulScalar(0.5)
	}
	unitDir := r.Direction.Unit()
	a := 0.5 * (unitDir.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(a))
}
