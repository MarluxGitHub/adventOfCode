package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"bufio"
	"fmt"
	"log"
	"math/big"
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
var day = 7

type Way struct {
	IsSplitter bool
	IsRay      bool
	IsStart    bool
}

var grid map[datastructures.Point]Way
var height int
var width int

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

func genGrid() {
	grid = make(map[datastructures.Point]Way)
	height = len(lines)
	width = 0
	if height > 0 {
		for _, l := range lines {
			if len(l) > width {
				width = len(l)
			}
		}
	}

	for y, line := range lines {

		for x, r := range line {
			pt := datastructures.Point{X: x, Y: y}
			var way Way
			switch r {
			case '^':
				way.IsSplitter = true
			case 'S':
				way.IsStart = true
				way.IsRay = true
			}
			grid[pt] = way
		}
	}
}

// Solve part 1
func Solve1() {
	genGrid()

	// propagate row by row (top to bottom) so newly created rays from splitters
	// in the current row will be processed in the next row iteration
	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			pt := datastructures.Point{X: x, Y: y}
			above := datastructures.Point{X: x, Y: y - 1}

			if grid[above].IsRay {
				if grid[pt].IsSplitter {
					// create rays to left and right if within bounds
					if x-1 >= 0 {
						left := datastructures.Point{X: x - 1, Y: y}
						tmp := grid[left]
						tmp.IsRay = true
						grid[left] = tmp
					}
					if x+1 < width {
						right := datastructures.Point{X: x + 1, Y: y}
						tmp := grid[right]
						tmp.IsRay = true
						grid[right] = tmp
					}
					result++
				} else {
					tmp := grid[pt]
					tmp.IsRay = true
					grid[pt] = tmp
				}
			}
		}
	}

}

// Solve part 2
func Solve2() {
	genGrid()

	// counts at each point (number of timelines that have a particle at that point in current state)
	counts := make(map[datastructures.Point]*big.Int)

	// seed starts
	for pt, w := range grid {
		if w.IsStart {
			counts[pt] = big.NewInt(1)
		}
	}

	// propagate row by row
	for y := 1; y < height; y++ {
		newCounts := make(map[datastructures.Point]*big.Int)

		for x := 0; x < width; x++ {
			above := datastructures.Point{X: x, Y: y - 1}
			c := counts[above]
			if c == nil || c.Sign() == 0 {
				continue
			}

			cur := datastructures.Point{X: x, Y: y}
			if grid[cur].IsSplitter {
				// distribute to left and right at same row
				if x-1 >= 0 {
					left := datastructures.Point{X: x - 1, Y: y}
					if newCounts[left] == nil {
						newCounts[left] = new(big.Int)
					}
					newCounts[left].Add(newCounts[left], c)
				}
				if x+1 < width {
					right := datastructures.Point{X: x + 1, Y: y}
					if newCounts[right] == nil {
						newCounts[right] = new(big.Int)
					}
					newCounts[right].Add(newCounts[right], c)
				}
			} else {
				// continue down into this cell
				if newCounts[cur] == nil {
					newCounts[cur] = new(big.Int)
				}
				newCounts[cur].Add(newCounts[cur], c)
			}
		}

		counts = newCounts
	}

	// sum all remaining counts — these are final timelines
	total := big.NewInt(0)
	for _, v := range counts {
		total.Add(total, v)
	}

	// try to fit into int for existing printing
	if total.BitLen() <= 62 {
		i64 := total.Int64()
		result = int(i64)
	} else {
		// overflow case: print directly and set result to 0
		println("2:" + total.String())
		result = 0
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
