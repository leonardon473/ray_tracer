package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	math "math"
	"os"
	r "ray_tracer"
)

func main() {
	// Image
	const aspectRadio = 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRadio)
	const samplesPerPixel = 100
	const maxDepth = 50

	// World
	R := math.Cos(r.Pi / 4)
	var world r.HittableList

	materialLeft := r.Lambertian{Albedo: r.Color{Z: 1}}
	materialRight := r.Lambertian{Albedo: r.Color{X: 1}}

	world.Add(r.Sphere{Center: r.Point3{X: -R, Z: -1}, Radius: R, MatPtr: materialLeft})
	world.Add(r.Sphere{Center: r.Point3{X: R, Z: -1}, Radius: R, MatPtr: materialRight})

	// Camera
	cam := r.NewCamera(90.0, aspectRadio)

	// Render
	img := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for i := 0; i < imageWidth; i++ {
		for j := (imageHeight - 1); j >= 0; j-- {
			pixelColor := r.Color{}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + r.RandomFloat64()) / (float64(imageWidth) - 1.0)
				v := (float64(j) + r.RandomFloat64()) / (float64(imageHeight) - 1.0)
				ray := cam.GetRay(u, v)
				pixelColor = pixelColor.Add(rayColor(ray, world, maxDepth))
			}
			writeColor(img, i, j, pixelColor, samplesPerPixel, imageHeight)
		}
	}

	if err := SaveImage("image.jpg", img); err != nil {
		log.Fatal(err)
	}
}

func rayColor(ray r.Ray, world r.Hittable, depth int) r.Color {
	var rec r.HitRecord

	if depth <= 0 {
		return r.Color{}
	}
	if world.Hit(ray, 0.001, r.Infinity, &rec) {
		var scattered r.Ray
		var attenuation r.Color
		if rec.MatPtr.Scatter(&ray, &rec, &attenuation, &scattered) {
			return rayColor(scattered, world, depth-1).Mul(attenuation)
		}
		return r.Color{}
	}
	unitDirection := ray.Direction.UnitVector()
	t := 0.5 * (unitDirection.Y + 1.0)
	return r.Color{X: 1.0, Y: 1.0, Z: 1.0}.Scale(1.0 - t).Add(r.Color{X: 0.5, Y: 0.7, Z: 1.0}.Scale(t))
}

func hitSphere(center r.Point3, radius float64, r r.Ray) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-halfB - math.Sqrt(discriminant)) / a
	}
}

func colorToNRGBA(rayColor r.Color) color.NRGBA {
	return color.NRGBA{
		R: uint8(int(255.999 * rayColor.X)),
		G: uint8(int(255.999 * rayColor.Y)),
		B: uint8(int(255.999 * rayColor.Z)),
		A: 255,
	}
}

func SaveImage(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	jpeg.Encode(f, img, &jpeg.Options{
		Quality: 95,
	})
	return nil
}

func mapValue(value, fromLow, fromHigh, toLow, toHigh float64) float64 {
	return (value-fromLow)*(toHigh-toLow)/(fromHigh-fromLow) + toLow
}

func writeColor(img *image.NRGBA, x, y int, pixelColor r.Color, samplesPerPixel int, imageHeight int) {
	red := pixelColor.X
	green := pixelColor.Y
	blue := pixelColor.Z

	scale := 1.0 / float64(samplesPerPixel)
	red = math.Sqrt(red * scale)
	green = math.Sqrt(green * scale)
	blue = math.Sqrt(blue * scale)

	pixelColor = r.Color{
		X: r.Clamp(red, 0.0, 0.999),
		Y: r.Clamp(green, 0.0, 0.999),
		Z: r.Clamp(blue, 0.0, 0.999),
	}
	yy := int(mapValue(float64(y), 0, float64(imageHeight), float64(imageHeight), 0))
	img.Set(x, yy, colorToNRGBA(pixelColor))
}
