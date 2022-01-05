package ray_tracer

type HittableList struct {
	objects []Hittable
}

func (h HittableList) clear() {
	h.objects = nil
}

func (h HittableList) add(object Hittable) {
	h.objects = append(h.objects, object)
}

func (h HittableList) hit(r Ray, tMin, tMax float64, rec HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, object := range h.objects {
		if object.Hit(r, tMin, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = tempRec
		}
	}

	return hitAnything
}
