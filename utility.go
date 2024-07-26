package main

import (
	"math"
	"math/rand"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func RandomFloatBetweenMinAndMax(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func RandomFloat() float64 {
	return rand.Float64()
}
