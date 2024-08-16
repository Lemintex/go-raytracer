package main

import "math"

type Ray struct {
	Origin    Vec3
	Direction Vec3
	Time      float64
}

func (r Ray) PointAt(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

func (r Ray) Color(world HittableList, c Camera, bounce int) Vec3 {
	if bounce == 0 {
		return Vec3{0, 0, 0}
	}
	didHit, hit := world.Hit(r, Interval{0.00001, math.Inf(1)})
	if !didHit {
		return c.Background
	}
	emitted := hit.Material.Emitted(hit.U, hit.V, hit.Point)

	didScatter, scattered, attenuation := hit.Material.Scatter(r, hit)
	if !didScatter {
		if emitted.X != 0 || emitted.Y != 0 || emitted.Z != 0 {
			return emitted
		}
	}
	scatterCol := attenuation.Mul(scattered.Color(world, c, bounce-1))
	return emitted.Add(scatterCol)
}

func GetRay(c Camera, x, y int) Ray {
	offset := SampleSquare()
	pixelSample := c.Pixel00Location.Add(c.PixelDeltaU.MulScalar(float64(x) + offset.X).Add(c.PixelDeltaV.MulScalar(float64(y) + offset.Y)))
	var rayOrigin Vec3
	if c.DefocusAngle <= 0 {
		rayOrigin = c.Origin
	} else {
		rayOrigin = DefocusDiskSample(c)
	}
	rayDirection := pixelSample.Sub(rayOrigin)
	time := RandomFloat()
	return Ray{rayOrigin, rayDirection, time}
}

func DefocusDiskSample(c Camera) Vec3 {
	p := RandomInUnitDisk()
	return c.Origin.Add(c.DefocusDiskU.MulScalar(p.X)).Add(c.DefocusDiskV.MulScalar(p.Y))
}
