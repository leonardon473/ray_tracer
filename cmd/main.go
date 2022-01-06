package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	r "ray_tracer"
)

func main() {
	// Image
	const aspectRadio = 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRadio)

	// World
	var world r.HittableList
	world.Add(r.Sphere{Center: r.Point3{Z: -1}, Radius: 0.5})
	world.Add(r.Sphere{Center: r.Point3{Y: -100.5, Z: -1}, Radius: 100})
	fmt.Println(world)
	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRadio * viewportHeight
	focalLength := 1.0

	origin := r.Point3{}
	horizontal := r.Vec3{X: viewportWidth}
	vertical := r.Vec3{Y: viewportHeight}
	lowerLeftCorner := origin.Sub(horizontal.Decrease(2)).Sub(vertical.Decrease(2)).Sub(r.Vec3{Z: focalLength})

	// Render
	img := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for i := 0; i < imageWidth; i++ {
		for j := (imageHeight - 1); j >= 0; j-- {
			u := float64(i) / (float64(imageWidth) - 1.0)
			v := float64(j) / (float64(imageHeight) - 1.0)
			direction := lowerLeftCorner.Add(horizontal.Scale(u)).Add(vertical.Scale(v)).Sub(origin)
			ray := r.Ray{Origin: origin, Direction: direction}
			pixelColor := rayColor(ray, world)
			jj := int(mapValue(float64(j), 0, float64(imageHeight), float64(imageHeight), 0))
			img.Set(i, jj, colorToNRGBA(pixelColor))
		}
	}

	if err := SaveImage("image.jpg", img); err != nil {
		log.Fatal(err)
	}
}

func rayColor(ray r.Ray, world r.Hittable) r.Color {
	var rec r.HitRecord
	if world.Hit(ray, 0, r.Infinity, &rec) {
		return rec.Normal.Add(r.Color{X: 1, Y: 1, Z: 1}).Scale(0.5)
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
