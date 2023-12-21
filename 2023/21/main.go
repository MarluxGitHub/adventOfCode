package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
	"gonum.org/v1/gonum/mat"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 21

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
	result = solve(64)
}

// Solve part 2
func Solve2() {
	a0 := solve(65)
	a1 := solve(65 + 131)
	a2 := solve(65 + 2*131)

	vandermonde := mat.NewDense(3, 3, []float64{0, 0, 1, 1, 1, 1, 4, 2, 1})
	b := mat.NewVecDense(3, []float64{float64(a0), float64(a1), float64(a2)})

	var x mat.VecDense
	err := x.SolveVec(vandermonde, b)
	if err != nil {
		log.Fatal(err)
	}

	n := 202300
	result = int(x.At(0, 0)*float64(n*n) + x.At(1, 0)*float64(n) + x.At(2, 0))

}

func solve(n int) int {
	if n > 64 {
		var newData []string
		for i := 0; i < 5; i++ {
			for _, line := range lines {
				newData = append(newData, strings.Repeat(strings.ReplaceAll(line, "S", "."), 5))
			}
		}
		lines = newData
	}

	width := len(lines[0])
	height := len(lines)

	type Point struct {
		x, y, steps int
	}

	q := list.New()
	sx, sy := width/2, height/2
	q.PushBack(Point{sx, sy, 0})

	s64 := make(map[Point]bool)
	visited := make(map[Point]bool)

	for q.Len() > 0 {
		e := q.Front()
		point := e.Value.(Point)
		q.Remove(e)

		if visited[point] {
			continue
		}
		visited[point] = true

		if point.steps == n {
			s64[Point{point.x, point.y, 0}] = true
		} else {
			if point.x >= 0 && lines[point.y][point.x-1] != '#' {
				q.PushBack(Point{point.x - 1, point.y, point.steps + 1})
			}
			if point.x < width-1 && lines[point.y][point.x+1] != '#' {
				q.PushBack(Point{point.x + 1, point.y, point.steps + 1})
			}
			if point.y >= 0 && lines[point.y-1][point.x] != '#' {
				q.PushBack(Point{point.x, point.y - 1, point.steps + 1})
			}
			if point.y < height-1 && lines[point.y+1][point.x] != '#' {
				q.PushBack(Point{point.x, point.y + 1, point.steps + 1})
			}
		}
	}

	if n == 64 {
		for y, line := range lines {
			ll := []rune(line)
			for point := range s64 {
				if point.y == y {
					ll[point.x] = 'O'
				}
			}
		}
	}

	return len(s64)
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
