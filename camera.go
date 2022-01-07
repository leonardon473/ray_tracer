package ray_tracer

import "math"

type Camera struct {
	Origin          Point3
	LowerLeftCorner Point3
	Horizontal      Vec3
	Vertical        Vec3
}

func NewCamera(vfov, aspectRatio float64) Camera {
	theta := DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	const focalLength = 1.0

	origin := Point3{X: 0, Y: 0, Z: 0}
	horizontal := Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := Vec3{X: 0, Y: viewportHeight, Z: 0}
	lowerLeftCorner := origin.Sub(horizontal.Decrease(2)).Sub(vertical.Decrease(2)).Sub(Vec3{X: 0, Y: 0, Z: focalLength})

	return Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
	}
}

func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Scale(u)).Add(c.Vertical.Scale(v)).Sub(c.Origin),
	}
}
