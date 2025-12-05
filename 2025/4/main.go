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
var year = 2025
var day = 4

var grid map[datastructures.Point]rune

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genGrid()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func genGrid() {
	grid = make(map[datastructures.Point]rune)
	for y, line := range lines {
		for x, r := range line {
			grid[datastructures.Point{X: x, Y: y}] = r
		}
	}
}

// Solve part 1
func Solve1() {
	adjacentDirections := []datastructures.Point{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: -1, Y: -1},
	}

	for point, value := range grid {
		if value == '.' {
			continue
		}

		rolls := 0

		for _, direction := range adjacentDirections {
			x := point.X + direction.X
			y := point.Y + direction.Y

			if grid[datastructures.Point{X: x, Y: y}] == '@' {
				rolls++
			}
		}

		if rolls < 4 {
			result++
		}
	}
}

func printGrid() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Print(string(grid[datastructures.Point{X: x, Y: y}]))
		}
		fmt.Println()
	}
}

// Solve part 2
func Solve2() {
	adjacentDirections := []datastructures.Point{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: -1, Y: -1},
	}

	removed := 1

	for removed > 0 {
		removed = 0

		for point, value := range grid {
			if value == '.' {
				continue
			}

			rolls := 0

			for _, direction := range adjacentDirections {
				x := point.X + direction.X
				y := point.Y + direction.Y

				if grid[datastructures.Point{X: x, Y: y}] == '@' || grid[datastructures.Point{X: x, Y: y}] == 'x' {
					rolls++
				}
			}

			if rolls < 4 {
				grid[point] = 'x'
			}
		}

		for point, value := range grid {
			if value == 'x' {
				grid[point] = '.'
				removed++
			}
		}

		result += removed
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
