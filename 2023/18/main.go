package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"MarluxGitHub/adventOfCode/pkg/math"
	gomath "math"

	"bufio"
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 18

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

var dirs = map[string]datastructures.Point{
	"R": {X: 1, Y: 0},
	"D": {X: 0, Y: -1},
	"L": {X: -1, Y: 0},
	"U": {X: 0, Y: 1},
}

// Solve part 1
func Solve1() {
	coords := []datastructures.Point{{0, 0}}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir, distStr := parts[0], parts[1]
		dist, _ := strconv.Atoi(distStr)

		last := coords[len(coords)-1]
		newPoint := datastructures.Point{last.X + dist*dirs[dir].X, last.Y + dist*dirs[dir].Y}
		coords = append(coords, newPoint)

	}

	bounds := 0
	for i := 0; i < len(coords)-1; i++ {
		bounds += math.Abs(coords[i+1].X-coords[i].X) + math.Abs(coords[i+1].Y-coords[i].Y)
	}

	area := polygonArea(coords)
	result = int(float64(area) - float64(bounds)/2 + 1 + float64(bounds))
}

// Solve part 2
func Solve2() {
	coords := []datastructures.Point{{0, 0}}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir, _, code := parts[0], parts[1], parts[2]

		dir = []string{"R", "D", "L", "U"}[int(code[len(code)-2]-'0')]
		dist64, _ := strconv.ParseInt(code[2:len(code)-2], 16, 64)
		dist := int(dist64)

		last := coords[len(coords)-1]
		newPoint := datastructures.Point{last.X + dist*dirs[dir].X, last.Y + dist*dirs[dir].Y}
		coords = append(coords, newPoint)

	}

	bounds := 0
	for i := 0; i < len(coords)-1; i++ {
		bounds += math.Abs(coords[i+1].X-coords[i].X) + math.Abs(coords[i+1].Y-coords[i].Y)
	}

	area := polygonArea(coords)
	result = int(float64(area) - float64(bounds)/2 + 1 + float64(bounds))
}

func polygonArea(coords []datastructures.Point) float64 {
	area := 0.0
	j := len(coords) - 1

	for i := 0; i < len(coords); i++ {
		area += (float64(coords[j].X+coords[i].X) * float64(coords[j].Y-coords[i].Y))
		j = i
	}

	return gomath.Abs(area / 2)
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
