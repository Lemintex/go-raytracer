package main

import (
	"math"
	"sync"
)

type Camera struct {
	SamplesPerPixel         int
	AspectRatio             float64
	FocalLength             float64
	ImageWidth              int
	ImageHeight             int
	viewportWidth           float64
	viewportHeight          float64
	ViewportFOV             int
	viewportUpperLeftCorner Vec3
	Origin                  Vec3
	Pixel00Location         Vec3
	ViewportU               Vec3
	ViewportV               Vec3
	PixelDeltaU             Vec3
	PixelDeltaV             Vec3

	// position of the camera
	LookFrom Vec3
	LookAt   Vec3
	VUp      Vec3
}

func (c *Camera) Render(image []ImageLine, world HittableList) []ImageLine {
	wg := sync.WaitGroup{}
	wg.Add(c.ImageHeight)
	for y := range c.ImageHeight {
		go c.ProcessLine(world, &image[y], &wg)
	}
	wg.Wait()
	return image
}
func (c *Camera) Initialize() {
	// camera position
	c.LookFrom = Vec3{-2, 2, 1}
	c.Origin = c.LookFrom
	c.LookAt = Vec3{0, 0, -1}
	c.VUp = Vec3{0, 1, 0}

	// image info
	c.AspectRatio = 16.0 / 9.0
	c.ImageWidth = 1920
	c.ImageHeight = int(float64(c.ImageWidth) / c.AspectRatio)

	// camera info
	c.SamplesPerPixel = 2
	c.FocalLength = c.LookFrom.Sub(c.LookAt).Length()

	// camera coordinate frame
	w := c.LookFrom.Sub(c.LookAt).Unit()
	u := c.VUp.Cross(w).Unit()
	v := w.Cross(u)
	// viewport info
	c.ViewportFOV = 90
	theta := DegreesToRadians(float64(c.ViewportFOV))
	vh := math.Tan(theta / 2.0)
	c.viewportHeight = 2.0 * vh * c.FocalLength
	c.viewportWidth = c.AspectRatio * c.viewportHeight
	c.ViewportU, c.ViewportV = u.MulScalar(c.viewportWidth), v.MulScalar(-c.viewportHeight)

	// pixel info
	c.PixelDeltaU, c.PixelDeltaV = c.ViewportU.DivScalar(float64(c.ImageWidth)), c.ViewportV.DivScalar(float64(c.ImageHeight))

	h, v := c.ViewportU.MulScalar(0.5), c.ViewportV.MulScalar(0.5)
	c.viewportUpperLeftCorner = c.Origin.Sub(h).Sub(v).Sub(w.MulScalar(c.FocalLength))
	temp := c.PixelDeltaU.Add(c.PixelDeltaV).MulScalar(0.5)
	c.Pixel00Location = c.viewportUpperLeftCorner.Add(temp)

}

func (c Camera) ProcessLine(world HittableList, line *ImageLine, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range c.ImageWidth {
		var pixelColor Vec3
		for range c.SamplesPerPixel {
			r := GetRay(c, x, line.LineNumber)
			pixelColor = pixelColor.Add(r.Color(world, 25))
		}
		color := pixelColor.DivScalar(float64(c.SamplesPerPixel))
		line.Pixels[x] = WriteColor(color.X, color.Y, color.Z)
	}
}
