package ray_tracer

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) Invert() Vec3 {
	return Vec3{
		-v.X,
		-v.Y,
		-v.Z,
	}
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{
		v.X + v2.X,
		v.Y + v2.Y,
		v.Z + v2.Z,
	}
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{
		v.X - v2.X,
		v.Y - v2.Y,
		v.Z - v2.Z,
	}
}

func (v Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{
		v.X * v2.X,
		v.Y * v2.Y,
		v.Z * v2.Z,
	}
}

func (v Vec3) Div(v2 Vec3) Vec3 {
	return Vec3{
		v.X / v2.X,
		v.Y / v2.Y,
		v.Z / v2.Z,
	}
}

func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{
		v.X * s,
		v.Y * s,
		v.Z * s,
	}
}

func (v Vec3) Decrease(s float64) Vec3 {
	return v.Scale(1 / s)
}

func (v Vec3) length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vec3) UnitVector() Vec3 {
	return v.Decrease(v.length())
}

func (v Vec3) IsNearZero() bool {
	const s = 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}
func MakeRandomVec3() Vec3 {
	return Vec3{X: RandomFloat64(), Y: RandomFloat64(), Z: RandomFloat64()}
}

func MakeRandomRangeVec3(min, max float64) Vec3 {
	return Vec3{
		X: RandomRangeFloat64(min, max),
		Y: RandomRangeFloat64(min, max),
		Z: RandomRangeFloat64(min, max),
	}
}

func RandomInUnitSphere() Vec3 {
	for {
		p := MakeRandomRangeVec3(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func RandomUnitVector() Vec3 {
	return RandomInUnitSphere().UnitVector()
}

func MakeRandomInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := RandomInUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	} else {
		return inUnitSphere.Invert()
	}
}

type Point3 = Vec3
type Color = Vec3
