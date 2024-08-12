package main

import (
	"image"
	"image/jpeg"
	"os"
)

func ReadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	image, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	return image, err
}

func GetPixel(image image.Image, x, y int) (r, g, b, a uint32) {
	return image.At(x, y).RGBA()
}
