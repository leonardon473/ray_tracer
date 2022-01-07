package ray_tracer

type Material interface {
	Scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	Albedo Color
}

func (l Lambertian) Scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := rec.Normal.Add(RandomUnitVector())

	if scatterDirection.IsNearZero() {
		scatterDirection = rec.Normal
	}

	scattered = &Ray{rec.Point, scatterDirection}
	attenuation = &l.Albedo
	return true
}

type Metal struct {
	Albedo Color
}

func (m Metal) Scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := rIn.Direction.UnitVector().Reflect(rec.Normal)
	scattered = &Ray{rec.Point, reflected}
	attenuation = &m.Albedo
	return scattered.Direction.Dot(rec.Normal) > 0
}
