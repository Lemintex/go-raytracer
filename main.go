package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Pixel struct {
	R, G, B int
}

type ImageLine struct {
	LineNumber int
	Pixels     []Pixel
}

const (
	width  = 256
	height = 256
)

func main() {
	start := time.Now()
	image := make([]ImageLine, height)
	for y := range height {
		image[y].LineNumber = y
		image[y].Pixels = make([]Pixel, width)
	}
	f, err := os.Create("images/image1.ppm")
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
		r := float64(x) / float64(width)
		g := float64(line.LineNumber) / float64(height)
		b := 0.2

		ir := int(255.99 * r)
		ig := int(255.99 * g)
		ib := int(255.99 * b)

		line.Pixels[x] = Pixel{ir, ig, ib}
	}
}
