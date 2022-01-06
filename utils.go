package ray_tracer

import "math"

var Infinity = math.Inf(1)

const Pi = math.Pi

func degreesToRadians(degrees float64) float64 {
	return degrees * Pi / 180.0
}
