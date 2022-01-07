package ray_tracer

type Material interface {
	Scatter(rIn Ray, rec HitRecord, attenuation *Color, scattered *Ray)
}

type Lambertian struct {
	albedo Color
}

func (l Lambertian) Scatter(rIn Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := rec.Normal.Add(RandomUnitVector())

	if scatterDirection.IsNearZero() {
		scatterDirection = rec.Normal
	}
	
	scattered = &Ray{rec.Point, scatterDirection}
	attenuation = &l.albedo
	return true
}
