package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/echojc/aocutil"

	"MarluxGitHub/adventOfCode/pkg/datastructures"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2024
var day = 14

var Width = 11
var Height = 7

type Robot struct {
	Pos, Vel datastructures.Point
}

var Robots []Robot

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
	genRobots()

	for range 100 {
		for _, r := range Robots {
			r.Move()
		}
	}

	getResultPart1()
}

func (r *Robot) Move() {
	r.Pos.X += r.Vel.X
	if r.Pos.X < 0 {
		r.Pos.X = Width - r.Pos.X
	}

	if r.Pos.X >= Width {
		r.Pos.X = r.Pos.X - Width
	}

	r.Pos.Y += r.Vel.Y
	if r.Pos.Y < 0 {
		r.Pos.Y = Height - r.Pos.Y
	}

	if r.Pos.Y >= Height {
		r.Pos.Y = r.Pos.Y - Height
	}
}

func getResultPart1() {
	topLeft, topRight, bottomleft, bottomRight := 0, 0, 0, 0

	for _, r := range Robots {
		if r.Pos.X < Width/2 && r.Pos.Y < Height/2 {
			topLeft++
			continue
		}
		if r.Pos.X > Width/2 && r.Pos.Y < Height/2 {
			topRight++
			continue
		}
		if r.Pos.X < Width/2 && r.Pos.Y > Height/2 {
			bottomleft++
			continue
		}
		if r.Pos.X > Width/2 && r.Pos.Y > Height/2 {
			bottomRight++
			continue
		}
	}

	result = topLeft * topRight * bottomleft * bottomRight
}

// Solve part 2
func Solve2() {

}

func genRobots() {
	Robots = make([]Robot, len(lines))

	for i, line := range lines {
		var r Robot
		// p=0,4 v=3,-3
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.Pos.X, &r.Pos.Y, &r.Vel.X, &r.Vel.Y)

		if err != nil {
			log.Fatal(err)
		}

		Robots[i] = r
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
