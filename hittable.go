package ray_tracer

type HitRecord struct {
	point     Point3
	Normal    Vec3
	t         float64
	frontFace bool
}

func (rec HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	frontFace := r.Direction.Dot(outwardNormal) < 0
	if frontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Invert()
	}

}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64, rec HitRecord) bool
}
