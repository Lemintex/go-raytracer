package main

import (
	"bufio"
	"fmt"
	"math"
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

	CreateScene()

	image := make([]ImageLine, cam.ImageHeight)
	for y := range cam.ImageHeight {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, cam.ImageWidth)
	}
	f, err := os.Create("images/Book 2/image2.ppm")
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

func CreateScene() {
	defer world.BuildBVH()
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			if math.Abs(float64(a)) <= 2 && math.Abs(float64(b)) <= 2 {
				continue
			}
			size := RandomFloatBetweenMinAndMax(0.1, 0.5)
			randomMat := RandomFloat()
			var material Material
			if randomMat < 0.8 {
				material = Lambertian{Vec3{RandomFloat(), RandomFloat(), RandomFloat()}}
			} else if randomMat < 0.95 {
				material = Metal{Vec3{0.5 * (1 + RandomFloat()), 0.5 * (1 + RandomFloat()), 0.5 * (1 + RandomFloat())}, 0.5 * RandomFloat()}
			} else {
				material = Dielectric{1.5}
			}
			center := Vec3{1.5*float64(a) + 0.9*RandomFloat(), size, 1.5*float64(b) + 0.9*RandomFloat()}
			world.Add(NewMovingSphere(center, center.Add(Vec3{0, 0.5, 0}), size, material))
		}
	}

	world.Add(NewStationarySphere(Vec3{0, -1000, 0}, 1000, Lambertian{Vec3{0.5, 0.5, 0.5}}))
	world.Add(NewStationarySphere(Vec3{0, .75, 0}, .75, Dielectric{1.5}))
	world.Add(NewStationarySphere(Vec3{0, .625, 0}, -0.5, Dielectric{1 / 1.5}))
	world.Add(NewStationarySphere(Vec3{1, .75, 1}, .75, Lambertian{Vec3{0.4, 0.2, 0.1}}))
	world.Add(NewStationarySphere(Vec3{-1, .75, -1}, .75, Metal{Vec3{0.7, 0.6, 0.5}, 0.0}))
}
