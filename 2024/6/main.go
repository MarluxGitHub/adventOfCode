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
var day = 6

var room map[datastructures.Point]rune
var pos datastructures.Point
var start datastructures.Point
var startDirection datastructures.Point
var direction datastructures.Point

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	readRoom()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {

	for {
		// mark current field as visited
		room[pos] = 'X'

		// check if the field of pos + direction is in room
		if _, ok := room[pos.Add(direction)]; !ok {
			break
		}

		// check if the field of pos + direction is a wall
		if room[pos.Add(direction)] == '#' {
			// Turn right
			if direction.X == 0 && direction.Y == -1 {
				direction = datastructures.Point{X: 1, Y: 0}
				continue
			}
			if direction.X == 1 && direction.Y == 0 {
				direction = datastructures.Point{X: 0, Y: 1}
				continue
			}
			if direction.X == 0 && direction.Y == 1 {
				direction = datastructures.Point{X: -1, Y: 0}
				continue
			}
			if direction.X == -1 && direction.Y == 0 {
				direction = datastructures.Point{X: 0, Y: -1}
				continue
			}

		}

		// move to the next field
		pos = pos.Add(direction)
	}

	// count the X in room
	for _, v := range room {
		if v == 'X' {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
	for r := range room {
		if room[r] == '.' {
			room[r] = '#'
		}
		for {
			// mark current field as visited
			if direction.X == 0 {
				// if already - then +
				if room[pos] == '-' {
					room[pos] = '+'
				} else {
					room[pos] = '|'
				}

			} else {
				// if already | then +
				if room[pos] == '|' {
					room[pos] = '+'
				} else {
					room[pos] = '-'
				}
			}

			// check if the field of pos + direction is in room
			if _, ok := room[pos.Add(direction)]; !ok {
				break
			}

			// check if field is already visited
			if (room[pos.Add(direction)] == '-' || room[pos.Add(direction)] == '+') && direction.X == 0 {
				result++
				break
			}

			if (room[pos.Add(direction)] == '|' || room[pos.Add(direction)] == '+') && direction.Y == 0 {
				result++
				break
			}

			// check if the field of pos + direction is a wall
			if room[pos.Add(direction)] == '#' {
				// Turn right
				if direction.X == 0 && direction.Y == -1 {
					direction = datastructures.Point{X: 1, Y: 0}
					continue
				}
				if direction.X == 1 && direction.Y == 0 {
					direction = datastructures.Point{X: 0, Y: 1}
					continue
				}
				if direction.X == 0 && direction.Y == 1 {
					direction = datastructures.Point{X: -1, Y: 0}
					continue
				}
				if direction.X == -1 && direction.Y == 0 {
					direction = datastructures.Point{X: 0, Y: -1}
					continue
				}

			}

			// move to the next field
			pos = pos.Add(direction)
		}
		cleanRoom()
		pos = start
	}

}

func readRoom() {
	room = make(map[datastructures.Point]rune)

	for y, line := range lines {
		for x, c := range line {
			room[datastructures.Point{x, y}] = c
			if c == '^' {
				pos = datastructures.Point{x, y}
				start = datastructures.Point{x, y}
				direction = datastructures.Point{0, -1}
				startDirection = datastructures.Point{0, -1}
			} else if c == 'v' {
				pos = datastructures.Point{x, y}
				start = datastructures.Point{x, y}
				direction = datastructures.Point{0, 1}
				startDirection = datastructures.Point{0, 1}
			} else if c == '<' {
				pos = datastructures.Point{x, y}
				start = datastructures.Point{x, y}
				direction = datastructures.Point{-1, 0}
				startDirection = datastructures.Point{-1, 0}
			} else if c == '>' {
				pos = datastructures.Point{x, y}
				start = datastructures.Point{x, y}
				direction = datastructures.Point{1, 0}
				startDirection = datastructures.Point{1, 0}
			}
		}
	}
}

func cleanRoom() {
	for k := range room {
		if room[k] == '|' || room[k] == '-' || room[k] == '+' || room[k] == 'X' {
			room[k] = '.'
		}
	}
}

func printRoom() {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if pos.X == x && pos.Y == y {
				printf("X")
			} else {
				printf(string(room[datastructures.Point{x, y}]))
			}
		}
		println("")
	}
	println("")
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
