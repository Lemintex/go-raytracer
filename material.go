package main

type Material interface {
	Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3)
}

type Lambertian struct {
	Albedo Vec3
}

func (l Lambertian) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	direction := hit.Normal.Add(RandomUnitVec3())
	if direction.IsNearZero() {
		direction = hit.Normal
	}
	scattered := Ray{hit.Point, direction}
	return true, scattered, l.Albedo
}

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m Metal) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	reflected := ray.Direction.Unit().Reflect(hit.Normal).Add(RandomVec3InUnitSphere().MulScalar(m.Fuzz))
	scattered := Ray{hit.Point, reflected}
	return scattered.Direction.Dot(hit.Normal) > 0, scattered, m.Albedo
}
