package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type ImageLine struct {
	LineNumber int
	Pixels     []Color
}

const (
	width       = 1920
	aspectRatio = 16.0 / 9.0
)

var pixel00Location Vec3
var cameraCenter Vec3
var pixelDeltaU Vec3
var pixelDeltaV Vec3
var world HittableList

func main() {
	start := time.Now()
	world = HittableList{}
	height := calcHeight()

	// camera
	cam := Camera{}

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	cameraCenter = Vec3{0, 0, 0}

	// Calculate teh vectors on the edges of the viewport
	viewportU, viewportV := Vec3{viewportWidth, 0, 0}, Vec3{0, -viewportHeight, 0}

	// Calculate the delta vectors between pixels
	pixelDeltaU, pixelDeltaV = viewportU.DivScalar(float64(width)), viewportV.DivScalar(float64(height))

	// Calculate the upper left corner of the viewport
	h, v := viewportU.MulScalar(.5), viewportV.MulScalar(.5)
	viewportUpperLeftCorner := cameraCenter.Sub(h).Sub(v).Sub(Vec3{0, 0, focalLength})
	temp := (pixelDeltaU.Add(pixelDeltaV)).MulScalar(0.5)
	pixel00Location = viewportUpperLeftCorner.Add(temp)

	// world
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5})
	world.Add(Sphere{Vec3{0, -100.5, -1}, 100})
	image := make([]ImageLine, height)
	for y := range height {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, width)
	}
	f, err := os.Create("images/image5.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "P3\n%d %d\n%d\n", width, height, 255)

	fmt.Println("P3")
	image = cam.Render(image, world)

	for y := range height {
		for x := range width {
			// fmt.Println(image[y].Pixels[x])
			fmt.Fprintf(w, "%d %d %d\n", image[y].Pixels[x].R, image[y].Pixels[x].G, image[y].Pixels[x].B)
		}
	}
	fmt.Println(time.Since(start))
}

func calcHeight() int {
	return int(float32(width) / aspectRatio)
}
