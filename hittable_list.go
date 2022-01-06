package ray_tracer

type HittableList struct {
	objects []Hittable
}

func (h HittableList) Clear() {
	h.objects = nil
}

func (h HittableList) Add(object Hittable) {
	h.objects = append(h.objects, object)
}

func (h HittableList) Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, object := range h.objects {
		if object.Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}
