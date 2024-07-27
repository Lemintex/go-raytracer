package main

import "math"

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) PointAt(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

func (r Ray) Color(world HittableList, bounce int) Vec3 {
	if bounce == 0 {
		return Vec3{0, 0, 0}
	}
	didHit, hit := world.Hit(r, Interval{0.001, math.Inf(1)})
	if didHit {
		direction := hit.Normal.Add(RandomUnitVec3())
		r := Ray{hit.Point, direction}
		return r.Color(world, bounce-1).MulScalar(0.5)
	}
	unitDir := r.Direction.Unit()
	a := 0.5 * (unitDir.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(a))
}

func GetRay(c Camera, x, y int) Ray {
	offset := SampleSquare()
	pixelSample := c.Pixel00Location.Add(c.PixelDeltaU.MulScalar(float64(x) + offset.X).Add(c.PixelDeltaV.MulScalar(float64(y) + offset.Y)))
	rayOrigin := c.Origin
	rayDirection := pixelSample.Sub(c.Origin)
	return Ray{rayOrigin, rayDirection}
}
