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

	*scattered = Ray{rec.Point, scatterDirection}
	*attenuation = l.Albedo
	return true
}

type Metal struct {
	Albedo Color
	Fuzz   float64
}

func MakeMetal(albedo Color, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{albedo, fuzz}
}

func (m Metal) Scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := rIn.Direction.UnitVector().Reflect(rec.Normal)
	*scattered = Ray{rec.Point, reflected.Add(RandomInUnitSphere().Scale(m.Fuzz))}
	*attenuation = m.Albedo
	return scattered.Direction.Dot(rec.Normal) > 0
}
