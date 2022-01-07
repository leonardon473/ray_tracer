package ray_tracer

import (
	"math"
	"math/rand"
)

var Infinity = math.Inf(1)

const Pi = math.Pi

func DegreesToRadians(degrees float64) float64 {
	return degrees * Pi / 180.0
}

func RandomFloat64() float64 {
	return rand.Float64()
}

func RandomRangeFloat64(min, max float64) float64 {
	return min + (max-min)*RandomFloat64()
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
