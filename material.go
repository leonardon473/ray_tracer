package ray_tracer

type Material interface {
	scatter(rIn Ray, rec HitRecord, attenuation Color, scattered Ray)
}
