package ray_tracer

import "math"

type Sphere struct {
	center Point3
	radius float64
}

func (s Sphere) Hit(r Ray, tMin, tMax float64, rec HitRecord) bool {
	oc := r.Origin.Sub(s.center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtD := math.Sqrt(discriminant)

	root := (-halfB - sqrtD) / a
	if root < tMin || tMax < root {
		root := (-halfB + sqrtD) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.T = root
	rec.Point = r.At(rec.T)
	outwardNormal := rec.Point.Sub(s.center).Decrease(s.radius)
	rec.SetFaceNormal(r, outwardNormal)
	return true
}
