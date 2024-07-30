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
	materialGround := Lambertian{Vec3{0.8, 0.8, 0.0}}
	materialCenter := Lambertian{Vec3{0.1, 0.2, 0.5}}
	materialLeft := Metal{Vec3{0.8, 0.8, 0.8}, 0.3}
	materialRight := Dielectric{1.5}
	materialRightBubble := Dielectric{1.0 / 1.5}

	world.Add(Sphere{Vec3{0, -100.5, -1}, 100, materialGround})
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5, materialCenter})
	world.Add(Sphere{Vec3{-1, 0, -1}, 0.5, materialLeft})
	world.Add(Sphere{Vec3{1, 0, -1}, 0.5, materialRight})
	world.Add(Sphere{Vec3{1, 0, -1}, 0.4, materialRightBubble})

	image := make([]ImageLine, cam.ImageHeight)
	for y := range cam.ImageHeight {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, cam.ImageWidth)
	}
	f, err := os.Create("images/image20.ppm")
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
