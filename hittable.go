package ray_tracer

type HitRecord struct {
	Point     Point3
	Normal    Vec3
	T         float64
	FrontFace bool
}

func (rec HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	rec.FrontFace = r.Direction.Dot(outwardNormal) < 0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Invert()
	}

}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool
}
