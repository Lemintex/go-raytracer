package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	width  = 256
	height = 256
)

func main() {
	start := time.Now()
	f, err := os.Create("images/output1.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "P3\n%d %d\n%d\n", width, height, 255)

	for y := range height {
		for x := range width {
			r := float64(x) / float64(width)
			g := float64(y) / float64(height)
			b := 0.2

			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)

			fmt.Fprintf(w, "%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Println(time.Since(start))
}
