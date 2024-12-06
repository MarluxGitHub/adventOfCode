package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

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
	cleanRoom()
	pos = start
	direction = startDirection
	res := 0
	var wg sync.WaitGroup
	i := 0

	for r := range room {
		// copy room map
		roomCopy := make(map[datastructures.Point]rune)
		for k, v := range room {
			roomCopy[k] = v
		}

		wg.Add(1)
		i++
		go processRoom(i, &wg, r, roomCopy, start, startDirection, &res)
	}
	wg.Wait()

	result = res
}

func processRoom(pNumber int, wg *sync.WaitGroup, r datastructures.Point, room map[datastructures.Point]rune, start datastructures.Point, startDirection datastructures.Point, result *int) {
	defer wg.Done()
	pos := start
	direction := startDirection
	visited := make(map[datastructures.Point]bool)
	i := 0

	if room[r] == '.' {
		room[r] = '#'
	} else {
		return
	}
	for {
		i++
		if i > 16900 {
			*result++
			break
		}

		// check if the field of pos + direction is in room
		nextPos := pos.Add(direction)
		if _, ok := room[nextPos]; !ok {
			break
		}

		// check if the field of pos + direction is a wall
		if room[nextPos] == '#' {
			// Turn right
			if direction.X == 0 && direction.Y == -1 {
				direction = datastructures.Point{X: 1, Y: 0}
			} else if direction.X == 1 && direction.Y == 0 {
				direction = datastructures.Point{X: 0, Y: 1}
			} else if direction.X == 0 && direction.Y == 1 {
				direction = datastructures.Point{X: -1, Y: 0}
			} else if direction.X == -1 && direction.Y == 0 {
				direction = datastructures.Point{X: 0, Y: -1}
			}
			continue
		}

		// move to the next field
		pos = nextPos
		visited[pos] = true
	}

	println("Debug:" + strconv.Itoa(pNumber) + " done")
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
		if room[k] != '#' {
			room[k] = '.'
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
