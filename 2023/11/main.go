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
var year = 2023
var day = 11

var space = [][]bool{}

var emptyVertical = map[int]bool{}
var emptyHorizontal = map[int]bool{}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseInput()
	expandVertical()
	expandHorizontal()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	solve(2)
}

// Solve part 2
func Solve2() {
	solve(1000000)
}

func solve(factor int) {
	// get Coordinates of all True Values
	galaxies := map[int]datastructures.Point{}

	i := 0
	for y := 0; y < len(space); y++ {
		for x := 0; x < len(space[y]); x++ {
			if space[y][x] {
				galaxies[i] = datastructures.Point{X: x, Y: y}
				i++
			}
		}
	}

	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {

			emptyHorizontalGalaxiesBetweenIAndJ := 0
			emptyVerticalGalaxiesBetweenIAndJ := 0

			startX := min(galaxy.X, galaxies[j].X) + 1
			endX := max(galaxy.X, galaxies[j].X)
			for k := startX; k < endX; k++ {
				if emptyHorizontal[k] {
					emptyHorizontalGalaxiesBetweenIAndJ++
				}
			}

			startY := min(galaxy.Y, galaxies[j].Y) + 1
			endY := max(galaxy.Y, galaxies[j].Y)
			for k := startY; k < endY; k++ {
				if emptyVertical[k] {
					emptyVerticalGalaxiesBetweenIAndJ++
				}
			}

			result += galaxy.ManhattanDistance(galaxies[j]) + (emptyHorizontalGalaxiesBetweenIAndJ * (factor - 1)) + (emptyVerticalGalaxiesBetweenIAndJ * (factor - 1))
		}
	}
}

func parseInput() {
	for y, line := range lines {
		space = append(space, []bool{})
		for _, c := range line {
			state := false
			if c == '#' {
				state = true
			}
			space[y] = append(space[y], state)
		}
	}
}

func expandVertical() {
	// look for lines without true values
	// if found, add a line below

	for y := 0; y < len(space); y++ {
		hasTrue := false
		for x := 0; x < len(space[y]); x++ {
			if space[y][x] {
				hasTrue = true
				break
			}
		}

		if !hasTrue {
			emptyVertical[y] = true
		}
	}
}

func expandHorizontal() {
	// look for rows without true values
	// if found, add a row to the right

	for x := 0; x < len(space[0]); x++ {
		hasTrue := false
		for y := 0; y < len(space); y++ {
			if space[y][x] {
				hasTrue = true
				break
			}
		}

		if !hasTrue {
			// add a row to the right of x
			emptyHorizontal[x] = true
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
