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
	width  = 256
	height = 256
)

func main() {
	start := time.Now()
	image := make([]ImageLine, height)
	for y := range height {
		image[y].LineNumber = y
		image[y].Pixels = make([]Color, width)
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
		r := float32(x) / float32(width)
		g := float32(line.LineNumber) / float32(height)
		b := float32(0.2)
		line.Pixels[x] = WriteColor(r, g, b)
	}
}
