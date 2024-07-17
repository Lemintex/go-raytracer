package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type ImageLine struct {
	LineNumber int
	Pixels     []Color
}

const (
	width  = 1920
	height = 1080
)

var pixel00Location Vec3
var cameraCenter Vec3
var pixelDeltaU Vec3
var pixelDeltaV Vec3

func main() {
	start := time.Now()
	height := calcHeight()

	// camera

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := (float64(width) / float64(height)) * viewportHeight
	cameraCenter = Vec3{0, 0, 0}

	viewport_u, viewport_v := Vec3{viewportWidth, 0, 0}, Vec3{0, viewportHeight, 0}
	pixelDeltaU, pixelDeltaV = viewport_u.DivScalar(float64(width)), viewport_v.DivScalar(float64(height))

	viewportLowerLeftCorner := cameraCenter.Sub(viewport_u.DivScalar(2)).Add(viewport_v.DivScalar(2)).Sub(Vec3{0, 0, focalLength})
	h, v := viewport_u.MulScalar(.5), viewport_v.MulScalar(.5)
	pixel00Location = viewportLowerLeftCorner.Sub(h.MulScalar(2).Add(v.MulScalar(2).Sub(Vec3{0, 0, focalLength})))
	image := make([]ImageLine, height)
	for y := range height {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, width)
	}
	f, err := os.Create("images/image2.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "P3\n%d %d\n%d\n", width, height, 255)

	wg := sync.WaitGroup{}
	wg.Add(height)
	for y := range height {
		go ProcessLine(&image[y], &wg)
	}
	wg.Wait()

	for _, line := range image {
		fmt.Println("First pixel: ", line.Pixels[0], "Last pixel: ", line.Pixels[len(line.Pixels)-1])
	}

	for y := range height {
		for x := range width {
			fmt.Fprintf(w, "%d %d %d\n", image[y].Pixels[x].R, image[y].Pixels[x].G, image[y].Pixels[x].B)
		}
	}
	fmt.Println(time.Since(start))
}

func ProcessLine(line *ImageLine, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range width {
		u, v := pixelDeltaU.MulScalar(float64(x)), pixelDeltaV.MulScalar(float64(line.LineNumber))
		pixelCenter := pixel00Location.Add(u).Add(v)
		rayDirection := pixelCenter.Sub(cameraCenter)
		r := Ray{cameraCenter, rayDirection}

		// color
		color := r.Color()
		line.Pixels[x] = WriteColor(color.X, color.Y, color.Z)
	}
}

func calcHeight() int {
	return int(float32(width) / 16.0 * 9.0)
}
