package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
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
var day = 17


func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	generateGrid()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	result = f(0,3)
}

// Solve part 2
func Solve2() {
	result = f(4,10)
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

type priorityQueue []*pathVertex

func (v *priorityQueue) Less(i, j int) bool {
	return (*v)[i].d < (*v)[j].d
}
func (v *priorityQueue) Len() int {
	return len(*v)
}
func (v *priorityQueue) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}
func (v *priorityQueue) Push(x interface{}) {
	*v = append(*v, x.(*pathVertex))
}
func (v *priorityQueue) Pop() interface{} {
	old := *v
	n := len(old)
	x := old[n-1]
	*v = old[0 : n-1]
	return x
}

// vertices that we follow via the priority queue and which form the path
type pathVertex struct {
	d      uint
	x, y   uint8
	dx, dy int8
}

// vertices that we mark as already visited. These are NOT just the vertices of the grid,
// so don't _just_ depend on the position in the grid, but also on the direction we came from.
type visitedVertex struct {
	x, y   uint8
	dx, dy int8
}

var grid [][]uint8

func generateGrid() {
	grid = make([][]uint8, len(lines))
	for i, line := range lines {
		grid[i] = make([]uint8, len(line))
		for j, c := range line {
			grid[i][j] = uint8(c - '0')
		}
	}
}

func f(MIN float64, MAX int8) int {
	w, h := uint8(len(grid[0])), uint8(len(grid))

	var q priorityQueue
	heap.Push(&q, &pathVertex{})
	var visited = make(map[visitedVertex]struct{})
	var c *pathVertex
	// we need to visit some more nodes
	for len(q) > 0 {
		c = heap.Pop(&q).(*pathVertex)
		// check if c vertex is the end and we reached it with the minimum amount of steps in the same direction
		if c.x == w-1 && c.y == h-1 && (math.Abs(float64(c.dx)) >= MIN || math.Abs(float64(c.dy)) >= MIN) {
			break
		}

		// check if we can go in any direction
		if c.x > 0 && (c.dx == 0 && c.dy == 0 || c.dx < 0 || math.Abs(float64(c.dy)) >= MIN) && c.dx > -MAX {
			v := visitedVertex{c.x - 1, c.y, c.dx - 1, 0}
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				heap.Push(&q, &pathVertex{c.d + uint(grid[c.y][c.x-1]), c.x - 1, c.y, c.dx - 1, 0})
			}
		}
		if c.x < w-1 && (c.dx == 0 && c.dy == 0 || c.dx > 0 || math.Abs(float64(c.dy)) >= MIN) && c.dx < MAX {
			v := visitedVertex{c.x + 1, c.y, c.dx + 1, 0}
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				heap.Push(&q, &pathVertex{c.d + uint(grid[c.y][c.x+1]), c.x + 1, c.y, c.dx + 1, 0})
			}
		}
		if c.y > 0 && (c.dx == 0 && c.dy == 0 || c.dy < 0 || math.Abs(float64(c.dx)) >= MIN) && c.dy > -MAX {
			v := visitedVertex{c.x, c.y - 1, 0, c.dy - 1}
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				heap.Push(&q, &pathVertex{c.d + uint(grid[c.y-1][c.x]), c.x, c.y - 1, 0, c.dy - 1})
			}
		}
		if c.y < h-1 && (c.dx == 0 && c.dy == 0 || c.dy > 0 || math.Abs(float64(c.dx)) >= MIN) && c.dy < MAX {
			v := visitedVertex{c.x, c.y + 1, 0, c.dy + 1}
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				heap.Push(&q, &pathVertex{c.d + uint(grid[c.y+1][c.x]), c.x, c.y + 1, 0, c.dy + 1})
			}
		}
	}

	return int(c.d)
}
