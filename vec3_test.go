package ray_tracer

import (
	"gotest.tools/assert"
	"testing"
)

func TestVec3CrossProduct(t *testing.T) {
	A := Vec3{3, -5, 4}
	B := Vec3{2, 6, 5}
	assert.Equal(t, A.Cross(B), Vec3{-49, -7, 28})
}

func TestVec3DotProduct(t *testing.T) {
	A := Vec3{3, -5, 4}
	B := Vec3{2, 6, 5}
	assert.Equal(t, A.Dot(B), -4.0)
}
