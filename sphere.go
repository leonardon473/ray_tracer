package ray_tracer

import "math"

type Sphere struct {
	Center Point3
	Radius float64
}

func (s Sphere) Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtD := math.Sqrt(discriminant)

	root := (-halfB - sqrtD) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtD) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.T = root
	rec.Point = r.At(rec.T)
	outwardNormal := rec.Point.Sub(s.Center).Decrease(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	return true
}
