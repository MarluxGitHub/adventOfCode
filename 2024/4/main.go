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
var day = 4

// map point to rune
var pointMap map[datastructures.Point]rune
var width, height int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genPointMap()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	// generate all possible strings of length 4 horizontally out of pointMap
	for x := 0; x < width-3; x++ {
		for y := 0; y < height; y++ {
			// check for xmas and samx
			if pointMap[datastructures.Point{x, y}] == 'X' && pointMap[datastructures.Point{x + 1, y}] == 'M' && pointMap[datastructures.Point{x + 2, y}] == 'A' && pointMap[datastructures.Point{x + 3, y}] == 'S' {
				result++
			}

			if pointMap[datastructures.Point{x, y}] == 'S' && pointMap[datastructures.Point{x + 1, y}] == 'A' && pointMap[datastructures.Point{x + 2, y}] == 'M' && pointMap[datastructures.Point{x + 3, y}] == 'X' {
				result++
			}
		}
	}

	// generate all possible strings of length 4 vertically out of pointMap
	for x := 0; x < width; x++ {
		for y := 0; y < height-3; y++ {
			// check for xmas and samx
			if pointMap[datastructures.Point{x, y}] == 'X' && pointMap[datastructures.Point{x, y + 1}] == 'M' && pointMap[datastructures.Point{x, y + 2}] == 'A' && pointMap[datastructures.Point{x, y + 3}] == 'S' {
				result++
			}

			if pointMap[datastructures.Point{x, y}] == 'S' && pointMap[datastructures.Point{x, y + 1}] == 'A' && pointMap[datastructures.Point{x, y + 2}] == 'M' && pointMap[datastructures.Point{x, y + 3}] == 'X' {
				result++
			}
		}
	}

	// generate all possible strings of length 4 diagonally (top-left to bottom-right) out of pointMap
	for x := 0; x < width-3; x++ {
		for y := 0; y < height-3; y++ {
			// check for xmas and samx
			if pointMap[datastructures.Point{x, y}] == 'X' && pointMap[datastructures.Point{x + 1, y + 1}] == 'M' && pointMap[datastructures.Point{x + 2, y + 2}] == 'A' && pointMap[datastructures.Point{x + 3, y + 3}] == 'S' {
				result++
			}

			if pointMap[datastructures.Point{x, y}] == 'S' && pointMap[datastructures.Point{x + 1, y + 1}] == 'A' && pointMap[datastructures.Point{x + 2, y + 2}] == 'M' && pointMap[datastructures.Point{x + 3, y + 3}] == 'X' {
				result++
			}
		}
	}

	// generate all possible strings of length 4 diagonally (bottom-left to top-right) out of pointMap
	for x := 0; x < width-3; x++ {
		for y := 3; y < height; y++ {
			// check for xmas and samx
			if pointMap[datastructures.Point{x, y}] == 'X' && pointMap[datastructures.Point{x + 1, y - 1}] == 'M' && pointMap[datastructures.Point{x + 2, y - 2}] == 'A' && pointMap[datastructures.Point{x + 3, y - 3}] == 'S' {
				result++
			}

			if pointMap[datastructures.Point{x, y}] == 'S' && pointMap[datastructures.Point{x + 1, y - 1}] == 'A' && pointMap[datastructures.Point{x + 2, y - 2}] == 'M' && pointMap[datastructures.Point{x + 3, y - 3}] == 'X' {
				result++
			}
		}
	}

}

// Solve part 2
func Solve2() {
	// go over wall pointmap and look for a
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if pointMap[datastructures.Point{X: x, Y: y}] == 'A' {
				if x-1 >= 0 && y-1 >= 0 && x+1 < width && y+1 < height {
					word1 := string(pointMap[datastructures.Point{X: x - 1, Y: y - 1}]) + "A" + string(pointMap[datastructures.Point{X: x + 1, Y: y + 1}])
					word2 := string(pointMap[datastructures.Point{X: x - 1, Y: y + 1}]) + "A" + string(pointMap[datastructures.Point{X: x + 1, Y: y - 1}])
					if ((word1 == "SAM") || (word1 == "MAS")) && ((word2 == "SAM") || (word2 == "MAS")) {
						result++
					}
				}

			}
		}
	}
}

func genPointMap() {
	pointMap = make(map[datastructures.Point]rune)

	height = len(lines)
	for x, line := range lines {
		width = len(line)
		for y, r := range line {
			pointMap[datastructures.Point{x, y}] = r
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
