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

var world HittableList

func main() {
	start := time.Now()
	world = HittableList{}

	// camera
	cam := Camera{}
	cam.Initialize()

	// world
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5})
	world.Add(Sphere{Vec3{0, -100.5, -1}, 100})
	image := make([]ImageLine, cam.ImageHeight)
	for y := range cam.ImageHeight {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, cam.ImageWidth)
	}
	f, err := os.Create("images/image8.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "P3\n%d %d\n%d\n", cam.ImageWidth, cam.ImageHeight, 255)

	image = cam.Render(image, world)

	for y := range cam.ImageHeight {
		for x := range cam.ImageWidth {
			fmt.Fprintf(w, "%d %d %d\n", image[y].Pixels[x].R, image[y].Pixels[x].G, image[y].Pixels[x].B)
		}
	}
	fmt.Println(time.Since(start))
}
