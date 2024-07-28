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
}

func (m Metal) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	reflected := ray.Direction.Unit().Reflect(hit.Normal)
	scattered := Ray{hit.Point, reflected}
	return true, scattered, m.Albedo
}
