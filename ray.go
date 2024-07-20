package main

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (v Vec3) At(t float64) Vec3 {
	return v.Add(v.MulScalar(t))
}

func (r Ray) Color() Vec3 {
	if r.HitSphere(Vec3{0, 0, -1}, 0.5) {
		return Vec3{1, 0, 0}
	}
	unitDir := r.Direction.Unit()
	a := 0.5 * (unitDir.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(a))
}

func (r Ray) HitSphere(center Vec3, radius float64) bool {
	oc := r.Origin.Sub(center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
