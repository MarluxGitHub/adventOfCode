package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
var year = 2024
var day = 10

var hiking = map[datastructures.Point]Lava{}

type Lava struct {
	value     int
	isReached bool
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genMap()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	for key, point := range hiking {
		if point.value == 0 {
			checkTrail1(key)
		}

		for k, p := range hiking {
			// reset all isReached
			p.isReached = false
			hiking[k] = p
		}
	}
}

// Solve part 2
func Solve2() {
	for key, point := range hiking {
		if point.value == 0 {
			checkTrail2(key)
		}
	}
}

func checkTrail1(point datastructures.Point) {
	if hiking[point].value == 9 {
		if !hiking[point].isReached {
			result++
		}

		lava := hiking[point]
		lava.isReached = true
		hiking[point] = lava

		return
	}

	// check top left bottom and right
	top := datastructures.Point{X: point.X, Y: point.Y - 1}
	left := datastructures.Point{X: point.X - 1, Y: point.Y}
	right := datastructures.Point{X: point.X + 1, Y: point.Y}
	bottom := datastructures.Point{X: point.X, Y: point.Y + 1}

	// loop over 4 directions
	for _, direction := range []datastructures.Point{top, left, right, bottom} {
		if hiking[direction].value == hiking[point].value+1 {
			checkTrail1(direction)
		}
	}
}

func checkTrail2(point datastructures.Point) {
	if hiking[point].value == 9 {
		result++
		return
	}

	// check top left bottom and right
	top := datastructures.Point{X: point.X, Y: point.Y - 1}
	left := datastructures.Point{X: point.X - 1, Y: point.Y}
	right := datastructures.Point{X: point.X + 1, Y: point.Y}
	bottom := datastructures.Point{X: point.X, Y: point.Y + 1}

	// loop over 4 directions
	for _, direction := range []datastructures.Point{top, left, right, bottom} {
		if hiking[direction].value == hiking[point].value+1 {
			checkTrail2(direction)
		}
	}
}

func genMap() {
	for y, line := range lines {
		for x, c := range line {
			hiking[datastructures.Point{X: x, Y: y}] = Lava{value: intval(c)}
		}
	}
}

func intval(c rune) int {
	val, err := strconv.Atoi(string(c))
	if err != nil {
		log.Fatal(err)
	}
	return val
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
