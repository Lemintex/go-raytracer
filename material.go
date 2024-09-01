package main

import "math"

type Material interface {
	Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3)
	Emitted(u, v float64, p Vec3) Vec3
}

type Lambertian struct {
	Albedo Vec3
	Tex    Texture
}

func (l Lambertian) Emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

func (l Lambertian) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	direction := hit.Normal.Add(RandomUnitVec3())
	if direction.IsNearZero() {
		direction = hit.Normal
	}
	scattered := Ray{hit.Point, direction, ray.Time}
	color := l.Albedo
	if l.Tex != nil {
		color = l.Tex.Value(hit.U, hit.V, hit.Point)
	}
	return true, scattered, color
}

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m Metal) Emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}
func (m Metal) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	reflected := ray.Direction.Unit().Reflect(hit.Normal).Add(RandomVec3InUnitSphere().MulScalar(m.Fuzz))
	scattered := Ray{hit.Point, reflected, ray.Time}
	return scattered.Direction.Dot(hit.Normal) > 0, scattered, m.Albedo
}

type Dielectric struct {
	RefractionIndex float64
}

func (d Dielectric) Emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

func (d Dielectric) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	attenuation := Vec3{1, 1, 1}
	ri := d.RefractionIndex
	if hit.FrontFace {
		ri = 1 / d.RefractionIndex
	}
	unitDirection := ray.Direction.Unit()

	cosTheta := math.Min(unitDirection.Neg().Dot(hit.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	var direction Vec3
	cannotRefract := ri*sinTheta > 1.0
	if cannotRefract || d.Schlick(cosTheta, ri) > RandomFloat() {
		direction = unitDirection.Reflect(hit.Normal)
	} else {
		direction = unitDirection.Refract(hit.Normal, ri)
	}

	scattered := Ray{hit.Point, direction, ray.Time}
	return true, scattered, attenuation
}

func (d Dielectric) Schlick(cosine float64, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

type DiffuseLight struct {
	Texture Texture
}

func (d DiffuseLight) Emitted(u, v float64, p Vec3) Vec3 {
	return d.Texture.Value(u, v, p)
}

func (d DiffuseLight) Scatter(ray Ray, hit HitInfo) (bool, Ray, Vec3) {
	return false, Ray{}, Vec3{0, 0, 0}
}
