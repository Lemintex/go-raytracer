package main

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (v Vec3) At(t float64) Vec3 {
	return v.Add(v.MulScalar(t))
}

func (r Ray) Color() Vec3 {
	unitDir := r.Direction.Unit()
	a := 0.5 * (unitDir.Y)
	return Vec3{0.1, 0.1, 0.1}.MulScalar(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(a))
}
