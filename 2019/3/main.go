package main

import (
	"MarluxGitHub/adventOfCode/pkg/math"
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
var year = 2019
var day = 3

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

// Solve part 1
func Solve1() {
	wire1 := strings.Split(lines[0], ",")
	wire2 := strings.Split(lines[1], ",")

	wire1Coords := make(map[string]bool)
	wire2Coords := make(map[string]bool)

	x, y := 0, 0

	for _, instruction := range wire1 {
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])

		for i := 0; i < distance; i++ {
			switch direction {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}

			wire1Coords[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}

	x, y = 0, 0

	for _, instruction := range wire2 {
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])

		for i := 0; i < distance; i++ {
			switch direction {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}

			wire2Coords[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}

	minDistance := 1000000

	for coord := range wire1Coords {
		coords := strings.Split(coord, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		if wire2Coords[coord] {
			distance := math.Abs(x) + math.Abs(y)
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	result = minDistance
}

// Solve part 2
func Solve2() {
	wire1 := strings.Split(lines[0], ",")
	wire2 := strings.Split(lines[1], ",")

	wire1Steps := make(map[string]int)
	wire2Steps := make(map[string]int)

	x, y, steps := 0, 0, 0

	for _, instruction := range wire1 {
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])

		for i := 0; i < distance; i++ {
			switch direction {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}

			steps++
			coord := fmt.Sprintf("%d,%d", x, y)
			if _, ok := wire1Steps[coord]; !ok {
				wire1Steps[coord] = steps
			}
		}
	}

	x, y, steps = 0, 0, 0

	for _, instruction := range wire2 {
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])

		for i := 0; i < distance; i++ {
			switch direction {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}

			steps++
			coord := fmt.Sprintf("%d,%d", x, y)
			if _, ok := wire2Steps[coord]; !ok {
				wire2Steps[coord] = steps
			}
		}
	}

	minSteps := 1000000

	for coord, steps1 := range wire1Steps {
		if steps2, ok := wire2Steps[coord]; ok {
			totalSteps := steps1 + steps2
			if totalSteps < minSteps {
				minSteps = totalSteps
			}
		}
	}

	result = minSteps
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
