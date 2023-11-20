package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
var year = 2016
var day = 1

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
	coord := datastructures.Point{0, 0}

	route := lines[0]
	direction := 0

	for _, step := range strings.Split(route, ", ") {
		if step[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
		}

		steps, err := strconv.Atoi(step[1:])

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 0:
			coord.Y += steps
		case 1:
			coord.X += steps
		case 2:
			coord.Y -= steps
		case 3:
			coord.X -= steps
		}
	}

	result = coord.ManhattanDistance(datastructures.Point{0, 0})
}

// Solve part 2
func Solve2() {
	coord := datastructures.Point{0, 0}

	route := lines[0]
	direction := 0

	locations := make(map[datastructures.Point]bool)

	for _, step := range strings.Split(route, ", ") {
		if step[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
		}

		steps, err := strconv.Atoi(step[1:])

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < steps; i++ {
			switch direction {
			case 0:
				coord.Y++
			case 1:
				coord.X++
			case 2:
				coord.Y--
			case 3:
				coord.X--
			}

			if locations[coord] {
				result = coord.ManhattanDistance(datastructures.Point{0, 0})
				return
			}

			locations[coord] = true
		}
	}

	result = coord.ManhattanDistance(datastructures.Point{0, 0})

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
