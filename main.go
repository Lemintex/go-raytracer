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
	width       = 1920
	aspectRatio = 16.0 / 9.0
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
	viewportWidth := aspectRatio * viewportHeight
	cameraCenter = Vec3{0, 0, 0}

	// Calculate teh vectors on the edges of the viewport
	viewportU, viewportV := Vec3{viewportWidth, 0, 0}, Vec3{0, viewportHeight, 0}

	// Calculate the delta vectors between pixels
	pixelDeltaU, pixelDeltaV = viewportU.DivScalar(float64(width)), viewportV.DivScalar(float64(height))

	// Calculate the upper left corner of the viewport
	h, v := viewportU.MulScalar(.5), viewportV.MulScalar(.5)
	viewportUpperLeftCorner := cameraCenter.Sub(h).Sub(v).Sub(Vec3{0, 0, focalLength})
	temp := (pixelDeltaU.Add(pixelDeltaV)).MulScalar(0.5)
	pixel00Location = viewportUpperLeftCorner.Add(temp)
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

	//	for _, line := range image {
	//		fmt.Println("First pixel: ", line.Pixels[0], "Last pixel: ", line.Pixels[len(line.Pixels)-1])
	//	}

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

		color := r.Color()
		line.Pixels[x] = WriteColor(color.X, color.Y, color.Z)
	}
}

func calcHeight() int {
	return int(float32(width) / aspectRatio)
}
