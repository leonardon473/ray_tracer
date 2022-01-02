package ray_tracer

type Ray struct {
	Origin    Point3
	Direction Vec3
}

func (r Ray) At(t float64) Point3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
