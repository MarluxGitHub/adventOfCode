package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 22

type Point struct {
	x, y, z int
}

type Brick struct {
	p, q Point
}

var Bricks []Brick

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	getBricks()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	for i := range Bricks {
		_, n := move(slices.Delete(slices.Clone(Bricks), i, i+1))
		if n == 0 {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
	for i := range Bricks {
		_, n := move(slices.Delete(slices.Clone(Bricks), i, i+1))
		result += n
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

func move(bricks []Brick) (b []Brick, m int) {
	bricks = slices.Clone(bricks)
loop:
	for i, a := range bricks {
		if a.p.z == 1 || a.q.z == 1 {
			continue
		}
		a.p.z, a.q.z = a.p.z-1, a.q.z-1

		for j, b := range bricks {
			if j != i && inter(a, b) {
				continue loop
			}
		}
		bricks[i] = a
		m++
	}
	return bricks, m
}

func inter(a, b Brick) bool {
	lint := func(a1, a2, b1, b2 int) bool {
		return min(a1, a2) <= max(b1, b2) && max(a1, a2) >= min(b1, b2)
	}
	return lint(a.p.x, a.q.x, b.p.x, b.q.x) &&
		lint(a.p.y, a.q.y, b.p.y, b.q.y) &&
		lint(a.p.z, a.q.z, b.p.z, b.q.z)
}

func getBricks() {
	for _, line := range lines {
		var b Brick
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d",
			&b.p.x, &b.p.y, &b.p.z, &b.q.x, &b.q.y, &b.q.z,
		)
		Bricks = append(Bricks, b)
	}
	slices.SortFunc(Bricks, func(a, b Brick) int {
		return cmp.Compare(min(a.p.z, a.q.z), min(b.p.z, b.q.z))
	})

	for n := 1; n != 0; Bricks, n = move(Bricks) {
	}
}
