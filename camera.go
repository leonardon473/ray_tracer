package ray_tracer

import "math"

type Camera struct {
	Origin          Point3
	LowerLeftCorner Point3
	Horizontal      Vec3
	Vertical        Vec3
	U, V, W         Vec3
	LensRadius      float64
}

func NewCamera(lookFrom, lookAt Point3, viewUp Vec3, vfov, aspectRatio, aperture, focusDisc float64) Camera {
	theta := DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.Sub(lookAt).UnitVector()
	u := viewUp.Cross(w).UnitVector()
	v := w.Cross(u)

	origin := lookFrom
	horizontal := u.Scale(focusDisc * viewportWidth)
	vertical := v.Scale(focusDisc * viewportHeight)
	lowerLeftCorner := origin.Sub(horizontal.Decrease(2)).
		Sub(vertical.Decrease(2)).Sub(w.Scale(focusDisc))

	lensRadius := aperture / 2

	return Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
		U:               u,
		V:               v,
		W:               w,
		LensRadius:      lensRadius,
	}
}

func (c Camera) GetRay(s, t float64) Ray {
	rd := MakeRandomInUnitDisk().Scale(c.LensRadius)
	offset := c.U.Scale(rd.X).Add(c.V.Scale(rd.Y))
	return Ray{
		Origin: c.Origin.Add(offset),
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Scale(s)).
			Add(c.Vertical.Scale(t)).Sub(c.Origin).Sub(offset),
	}
}
