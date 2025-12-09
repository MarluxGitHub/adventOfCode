package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"MarluxGitHub/adventOfCode/pkg/math"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 9

var Corners []datastructures.Point

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genCorners()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func genCorners() {
	Corners = make([]datastructures.Point, len(lines))

	for i, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		Corners[i] = datastructures.Point{X: x, Y: y}
	}
}

// Solve part 1
func Solve1() {
	n := len(Corners)
	if n < 2 {
		result = 0
		return
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		xi := Corners[i].X
		yi := Corners[i].Y
		for j := i + 1; j < n; j++ {
			xj := Corners[j].X
			yj := Corners[j].Y
			if xi == xj || yi == yj {
				// need opposite corners (different x and y)
				continue
			}
			area := (math.Abs(xi-xj) + 1) * (math.Abs(yi-yj) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	result = maxArea
}

// Helpers translated from the reference Python solution
func isPtOnSeg(px, py, x1, y1, x2, y2 int) bool {
	cross := (x2-x1)*(py-y1) - (y2-y1)*(px-x1)
	if cross != 0 {
		return false
	}
	if px < min(x1, x2) || px > max(x1, x2) {
		return false
	}
	if py < min(y1, y2) || py > max(y1, y2) {
		return false
	}
	return true
}

func isPtInPoly(px, py int, poly []datastructures.Point) bool {
	inside := false
	n := len(poly)
	for i := 0; i < n; i++ {
		x1 := poly[i].X
		y1 := poly[i].Y
		x2 := poly[(i+1)%n].X
		y2 := poly[(i+1)%n].Y
		if isPtOnSeg(px, py, x1, y1, x2, y2) {
			return true
		}
		if (y1 > py) != (y2 > py) {
			xIntersect := float64(x2-x1)*float64(py-y1)/float64(y2-y1) + float64(x1)
			if float64(px) < xIntersect {
				inside = !inside
			}
		}
	}
	return inside
}

func getOrient(ax, ay, bx, by, cx, cy int) int {
	v := (bx-ax)*(cy-ay) - (by-ay)*(cx-ax)
	if v > 0 {
		return 1
	}
	if v < 0 {
		return -1
	}
	return 0
}

func getSegInter(a [4]int, b [4]int) bool {
	o1 := getOrient(a[0], a[1], a[2], a[3], b[0], b[1])
	o2 := getOrient(a[0], a[1], a[2], a[3], b[2], b[3])
	o3 := getOrient(b[0], b[1], b[2], b[3], a[0], a[1])
	o4 := getOrient(b[0], b[1], b[2], b[3], a[2], a[3])
	return o1*o2 < 0 && o3*o4 < 0
}

func rectInPoly(x1, x2, y1, y2 int, points []datastructures.Point) bool {
	// check corners inside or on boundary
	corners := [][2]int{{x1, y1}, {x1, y2}, {x2, y1}, {x2, y2}}
	for _, c := range corners {
		if !isPtInPoly(c[0], c[1], points) {
			return false
		}
	}

	// check edges do not intersect polygon edges
	n := len(points)
	rectEdges := [][4]int{{x1, y1, x2, y1}, {x2, y1, x2, y2}, {x2, y2, x1, y2}, {x1, y2, x1, y1}}
	for _, re := range rectEdges {
		for i := 0; i < n; i++ {
			e2x1 := points[i].X
			e2y1 := points[i].Y
			e2x2 := points[(i+1)%n].X
			e2y2 := points[(i+1)%n].Y
			if getSegInter([4]int{re[0], re[1], re[2], re[3]}, [4]int{e2x1, e2y1, e2x2, e2y2}) {
				return false
			}
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Solve part 2
func Solve2() {
	// Use polygon containment + edge intersection test per reference Python solution
	n := len(Corners)
	if n < 2 {
		result = 0
		return
	}

	best := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := Corners[i]
			b := Corners[j]
			x1 := min(a.X, b.X)
			x2 := max(a.X, b.X)
			y1 := min(a.Y, b.Y)
			y2 := max(a.Y, b.Y)
			if x1 == x2 || y1 == y2 {
				continue
			}
			if rectInPoly(x1, x2, y1, y2, Corners) {
				area := (x2 - x1 + 1) * (y2 - y1 + 1)
				if area > best {
					best = area
				}
			}
		}
	}
	result = best
}

func readInput() {
	i, err := aocutil.NewInputFromFile("../../session_id")

	if err != nil {
		log.Fatal(err)
	}

	lines, err = i.Strings(year, day)

	if err != nil {
		log.Fatal(err)
	}
}
