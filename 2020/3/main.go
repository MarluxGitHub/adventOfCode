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
var year = 2020
var day = 3

var m map[datastructures.Point]byte
var width, height int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	readMapData()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	right, down := 3, 1

	result = sloping(right, down)
}

// Solve part 2
func Solve2() {

	result = sloping(1, 1) *
		sloping(3, 1) *
		sloping(5, 1) *
		sloping(7, 1) *
		sloping(1, 2)
}

func sloping(right, down int) int {
	x, y := right, down
	r := 0
	for y < height {
		if m[datastructures.Point{X: x, Y: y}] == '#' {
			r++
		}
		x = (x + right) % width
		y = y + down
	}
	return r
}

func readMapData() {
	height = len(lines)
	width = len(lines[0])
	m = make(map[datastructures.Point]byte)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			m[datastructures.Point{X: x, Y: y}] = lines[y][x]
		}
	}
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
