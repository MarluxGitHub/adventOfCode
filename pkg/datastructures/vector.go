package datastructures

import (
	"MarluxGitHub/adventOfCode/pkg/math"
	goMath "math"
	"strconv"
)

type Vector struct {
	X, Y, Z int
}

func (v Vector) Add(u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vector) Subtract(u Vector) Vector {
	return Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vector) Normalize() Vector {
	return Vector{math.Sign(v.X), math.Sign(v.Y), math.Sign(v.Z)}
}

func (v Vector) ManhattanDistance(u Vector) int {
	return math.Abs(v.X-u.X) + math.Abs(v.Y-u.Y) + math.Abs(v.Z-u.Z)
}

func (v Vector) EuclideanDistance(u Vector) float64 {
	dx := float64(v.X - u.X)
	dy := float64(v.Y - u.Y)
	dz := float64(v.Z - u.Z)
	return goMath.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (v Vector) ToString() string {
	return "(" + strconv.Itoa(v.X) + "," + strconv.Itoa(v.Y) + "," + strconv.Itoa(v.Z) + ")"
}
