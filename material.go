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

type Dielectric struct {
	Ir float64
}

func (d Dielectric) Scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool {
	*attenuation = Color{X: 1.0, Y: 1.0, Z: 1.0}
	var refractionRadio float64
	if rec.FrontFace {
		refractionRadio = 1.0 / d.Ir
	} else {
		refractionRadio = d.Ir
	}
	unitDirection := rIn.Direction.UnitVector()
	refracted := unitDirection.Refract(rec.Normal, refractionRadio)
	*scattered = Ray{rec.Point, refracted}
	return true

}
