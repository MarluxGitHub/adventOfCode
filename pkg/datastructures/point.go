package datastructures

import (
	"MarluxGitHub/adventOfCode/pkg/math"
	"strconv"
)

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Subtract(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Normalize() Point {
	return Point{math.Sign(p.X), math.Sign(p.Y)}
}

func (p Point) ManhattanDistance(q Point) int {
	return math.Abs(p.X-q.X) + math.Abs(p.Y-q.Y)
}

func (p Point) MooreDistance(q Point) int {
	return math.Max(math.Abs(p.X-q.X), math.Abs(p.Y-q.Y))
}

func (p Point) ToString() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}
