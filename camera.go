package main

import (
	"fmt"
)

type Camera struct {
	AspectRatio             float64
	FocalLength             float64
	ImageWidth              int
	Imageheight             int
	viewportWidth           float64
	viewportHeight          float64
	viewportUpperLeftCorner Vec3
	Origin                  Vec3
	Pixel00Location         Vec3
	ViewportU               Vec3
	ViewportV               Vec3
	PixelDeltaU             Vec3
	PixelDeltaV             Vec3
}

func (c *Camera) Render(image []ImageLine, world HittableList) []ImageLine {
	c.initialize()

	// wg := sync.WaitGroup{}
	// wg.Add(c.Imageheight)
	for y := range c.Imageheight {
		/*go*/ image[y] = c.ProcessLine(world, image[y]) //, &wg)
		fmt.Println("Line", y, "done")
		fmt.Println(image[y])
	}
	fmt.Println(world.Objects)
	// wg.Wait()
	return image
}
func (c *Camera) initialize() {
	c.AspectRatio = 16.0 / 9.0
	c.FocalLength = 1.0
	c.ImageWidth = 1920
	c.Imageheight = int(float64(c.ImageWidth) / c.AspectRatio)
	c.Origin = Vec3{0.0, 0.0, 0.0}
	c.viewportWidth = 2.0
	c.viewportHeight = c.viewportWidth / c.AspectRatio
	c.ViewportU = Vec3{c.viewportWidth, 0, 0}
	c.ViewportV = Vec3{0, -c.viewportHeight, 0}
	h, v := c.ViewportU.MulScalar(0.5), c.ViewportV.MulScalar(0.5)
	c.viewportUpperLeftCorner = c.Origin.Sub(h).Sub(v).Sub(Vec3{0, 0, c.FocalLength})
	temp := c.PixelDeltaU.Add(c.PixelDeltaV).MulScalar(0.5)
	c.Pixel00Location = c.viewportUpperLeftCorner.Add(temp)
}

func (c Camera) ProcessLine(world HittableList, line ImageLine /*, wg *sync.WaitGroup*/) ImageLine {
	// defer wg.Done()
	for x := range c.ImageWidth {
		u, v := c.PixelDeltaU.MulScalar(float64(x)), c.PixelDeltaV.MulScalar(float64(line.LineNumber))
		pixelCenter := c.Pixel00Location.Add(u).Add(v)
		rayDirection := pixelCenter.Sub(c.Origin)
		r := Ray{c.Origin, rayDirection}
		color := r.Color(world)
		line.Pixels[x] = Color{int(color.X * 255.999), int(color.Y * 255.999), int(color.Z * 255.999)}
	}
	return line
}
