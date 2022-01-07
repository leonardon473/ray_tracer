package ray_tracer

import "math"

type Camera struct {
	Origin          Point3
	LowerLeftCorner Point3
	Horizontal      Vec3
	Vertical        Vec3
}

func NewCamera(lookFrom, lookAt Point3, viewUp Vec3, vfov, aspectRatio float64) Camera {
	theta := DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.Sub(lookAt).UnitVector()
	u := viewUp.Cross(w).UnitVector()
	v := w.Cross(u)

	origin := lookFrom
	horizontal := u.Scale(viewportWidth)
	vertical := v.Scale(viewportHeight)
	lowerLeftCorner := origin.Sub(horizontal.Decrease(2)).Sub(vertical.Decrease(2)).Sub(w)

	return Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
	}
}

func (c Camera) GetRay(s, t float64) Ray {
	return Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Scale(s)).Add(c.Vertical.Scale(t)).Sub(c.Origin),
	}
}
