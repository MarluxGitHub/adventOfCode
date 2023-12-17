package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 16


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
	grid := build_light_grid()

	grid.illuminate_grid(0, 0, 1)

	result = grid.count_energized()
}

// Solve part 2
func Solve2() {
	max := 0
	grid := build_light_grid()

	for i := 0; i < grid.height; i++ {
		grid.illuminate_grid(i, 0, 1)
		max = int(math.Max(float64(max), float64(grid.count_energized())))
		grid.clear_energized()
		grid.illuminate_grid(i, grid.width-1, 3)
		max = int(math.Max(float64(max), float64(grid.count_energized())))
		grid.clear_energized()
	}
	for i := 0; i < grid.width; i++ {
		grid.illuminate_grid(0, i, 2)
		max = int(math.Max(float64(max), float64(grid.count_energized())))
		grid.clear_energized()
		grid.illuminate_grid(grid.height-1, i, 0)
		max = int(math.Max(float64(max), float64(grid.count_energized())))
		grid.clear_energized()
	}

	result = max
}


// 0 is north, 1 is east, 2 is south, 3 is west
var direction_table = map[int][]int{0: {-1, 0}, 1: {0, 1}, 2: {1, 0}, 3: {0, -1}}

func step_forward(x, y, dir, mirror int) (newx, newy, newdir int) {
	if mirror == 0 {
		// mirror = . or splitter that's ignored
		newdir = dir
	} else if mirror == 1 {
		// mirror = /
		newdir = ((1-dir)%4 + 4) % 4
	} else if mirror == 2 {
		// mirror = \
		newdir = ((3-dir)%4 + 4) % 4
	} else {
		panic("invalid direction for stepping")
	}
	// fmt.Println("new direction", newdir)
	newx = x + direction_table[newdir][0]
	newy = y + direction_table[newdir][1]
	return
}

type light_grid struct {
	mirrors   [][]int
	energized [][]bool
	height    int
	width     int
}

func (l light_grid) illuminate_grid(startx, starty, startdir int) {
	for i, j, dir := startx, starty, startdir; i >= 0 && j >= 0 && i < l.height && j < l.width; {
		curr_mirror := l.mirrors[i][j]
		//3 and 4 are splitters, - and | respectively
		if curr_mirror == 3 || curr_mirror == 4 {
			// fmt.Println("encountered splitter")
			if l.energized[i][j] {
				return
			} else {
				l.energized[i][j] = true
				if dir%2 == 1 {
					//left or right
					if curr_mirror == 3 {
						// continue going
						i, j, dir = step_forward(i, j, dir, 0)
					} else {
						//split
						l.illuminate_grid(i-1, j, 0)
						l.illuminate_grid(i+1, j, 2)
						return
					}
				} else {
					//up or down
					if curr_mirror == 4 {
						//continue going
						i, j, dir = step_forward(i, j, dir, 0)
					} else {
						//split
						l.illuminate_grid(i, j+1, 1)
						l.illuminate_grid(i, j-1, 3)
						return
					}
				}
			}
		} else {
			// fmt.Println("not a splitter", i, j, dir)
			l.energized[i][j] = true
			i, j, dir = step_forward(i, j, dir, curr_mirror)
		}
	}
	return
}

func build_light_grid() (grid light_grid) {
	for _, line := range lines {
		mirror_row := []int{}
		bool_row := []bool{}
		for _, char := range line {
			bool_row = append(bool_row, false)
			if char == '.' {
				mirror_row = append(mirror_row, 0)
			} else if char == '/' {
				mirror_row = append(mirror_row, 1)
			} else if char == '\\' {
				mirror_row = append(mirror_row, 2)
			} else if char == '-' {
				mirror_row = append(mirror_row, 3)
			} else if char == '|' {
				mirror_row = append(mirror_row, 4)
			} else {
				panic("invalid input during parsing")
			}
		}
		grid.mirrors = append(grid.mirrors, mirror_row)
		grid.energized = append(grid.energized, bool_row)
	}

	grid.height = len(grid.energized)
	grid.width = len(grid.energized[0])

	return
}

func (grid light_grid) count_energized() int {
	sum := 0
	for _, row := range grid.energized {
		for _, val := range row {
			if val {
				sum++
			}
		}
	}
	return sum
}

func (grid light_grid) clear_energized() {
	for i := 0; i < grid.height; i++ {
		for j := 0; j < grid.width; j++ {
			grid.energized[i][j] = false
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
