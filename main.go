package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

type ImageLine struct {
	LineNumber int
	Pixels     []Color
}

const SCENE_COUNT = 4

var world HittableList

func main() {
	start := time.Now()
	world = HittableList{}

	// flags
	scene := flag.Int("scene", 1, "scene id")
	flag.Parse()
	if *scene < 1 || *scene > SCENE_COUNT {
		*scene = 1
	}

	// scene
	CreateScene(scene)

	// camera
	cam := Camera{}
	cam.SetupCameraForScene(*scene)

	image := make([]ImageLine, cam.ImageHeight)
	for y := range cam.ImageHeight {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, cam.ImageWidth)
	}
	f, err := os.Create("images/Book 2/image7.ppm")
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

func CreateScene(scene *int) {
	defer world.BuildBVH()
	switch *scene {
	case 1:
		BouncingSpheres()

	case 2:
		CheckeredSpheres()

	case 3:
		Earth()

	case 4:
		PerlinNoise()
	}
}

func BouncingSpheres() {
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			if math.Abs(float64(a)) <= 2 && math.Abs(float64(b)) <= 2 {
				continue
			}
			size := RandomFloatBetweenMinAndMax(0.1, 0.5)
			randomMat := RandomFloat()
			var material Material
			if randomMat < 0.8 {
				material = Lambertian{Vec3{RandomFloat(), RandomFloat(), RandomFloat()}, nil}
			} else if randomMat < 0.95 {
				material = Metal{Vec3{0.5 * (1 + RandomFloat()), 0.5 * (1 + RandomFloat()), 0.5 * (1 + RandomFloat())}, 0.5 * RandomFloat()}
			} else {
				material = Dielectric{1.5}
			}
			center := Vec3{1.5*float64(a) + 0.9*RandomFloat(), size, 1.5*float64(b) + 0.9*RandomFloat()}
			world.Add(NewMovingSphere(center, center.Add(Vec3{0, 0.5, 0}), size, material))
		}
	}

	world.Add(NewStationarySphere(Vec3{0, -1000, 0}, 1000, Lambertian{Vec3{0.5, 0.5, 0.5}, CheckerTexture{0.32, SolidColor{Vec3{0.2, 0.3, 0.1}}, SolidColor{Vec3{1, 1, 1}}}}))
	world.Add(NewStationarySphere(Vec3{0, .75, 0}, .75, Dielectric{1.5}))
	world.Add(NewStationarySphere(Vec3{0, .625, 0}, -0.5, Dielectric{1 / 1.5}))
	world.Add(NewStationarySphere(Vec3{1, .75, 1}, .75, Lambertian{Vec3{0.4, 0.2, 0.1}, nil}))
	world.Add(NewStationarySphere(Vec3{-1, .75, -1}, .75, Metal{Vec3{0.7, 0.6, 0.5}, 0.0}))
}

// temporary
func CheckeredSpheres() {
	checker := CheckerTexture{0.2, SolidColor{Vec3{0.2, 0.3, 0.1}}, SolidColor{Vec3{0.9, 0.9, 0.9}}}

	world.Add(NewStationarySphere(Vec3{0, -1000, 0}, 1000, Lambertian{Vec3{0.5, 0.5, 0.5}, CheckerTexture{0.2, SolidColor{Vec3{0.6, 0.2, 0.3}}, SolidColor{Vec3{1, 1, 1}}}}))
	world.Add(NewStationarySphere(Vec3{0, .75, 0}, .75, Dielectric{1.5}))
	world.Add(NewStationarySphere(Vec3{0, .625, 0}, -0.5, Dielectric{1 / 1.5}))
	world.Add(NewStationarySphere(Vec3{1, .75, 1}, .75, Lambertian{Vec3{0.4, 0.2, 0.1}, checker}))
	world.Add(NewStationarySphere(Vec3{-1, .75, -1}, .75, Metal{Vec3{0.7, 0.6, 0.5}, 0.0}))
}

func Earth() {
	img, err := ReadImage("images/input/earthmap.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	world.Add(NewStationarySphere(Vec3{0, 0, 0}, 2, Lambertian{Vec3{1, 1, 1}, ImageTexture{img}}))
}

func PerlinNoise() {
	perlin := NewPerlin()
	world.Add(NewStationarySphere(Vec3{0, -1000, 0}, 1000, Lambertian{Vec3{0.5, 0.5, 0.5}, NoiseTexture{perlin, 7}}))
	world.Add(NewStationarySphere(Vec3{0, 2, 0}, 2, Lambertian{Vec3{0.5, 0.5, 0.5}, NoiseTexture{NewPerlin(), 2}}))

	world.Add(NewStationarySphere(Vec3{0, .625, 0}, -0.5, Dielectric{1 / 1.5}))
}
