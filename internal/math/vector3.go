package math

import "math"

type Vector3 struct {
	X, Y, Z int
}

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v Vector3) DstSquared(other Vector3) int {
	return squared(v.X-other.X) + squared(v.Y-other.Y) + squared(v.Z-other.Z)
}

func squared(a int) int {
	return int(math.Pow(float64(a), 2.0))
}
