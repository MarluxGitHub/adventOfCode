package main

import (
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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var oceanMap [][]int

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()
  initOceanMap()
  computeOceanMap()
  //printOceanMap()
  result := countCriticalPoints()
  println(strconv.Itoa(result))
}

func printOceanMap() {
	println("")
	for _, row := range oceanMap {
		println("")
		for _, col := range row {
			printf(strconv.Itoa(col))
		}
	}
	println("")
}

func initOceanMap() {
	oceanMap = make([][]int, 1000)

	for i := range oceanMap {
		oceanMap[i] = make([]int, 1000)
	}
}

func countCriticalPoints() int {
	count := 0
	for _, row := range oceanMap {
		for _, col := range row {
			if col > 1 {
				count++
			}
		}
	}
	return count
}

func computeOceanMap() {
	for _, line := range lines {
		params := strings.Split(line, "->")
		cordsStr := strings.Split(params[0], ",")
		cordGoal := strings.Split(params[1], ",")
		x1, y1 := cordsToInt(cordsStr)
		x2, y2 := cordsToInt(cordGoal)

		if(x1 == x2) {
			drawVerticalLine(x1, y1, y2)
			continue
		}

		if(y1 == y2) {
			drawHorizontalLine(x1, x2, y1)
			continue
		}

		drawDiagonalLine(x1, y1, x2, y2)
	}
}

func drawHorizontalLine(x1, x2, y int) {
	if(x1 > x2) {
		x1, x2 = x2, x1
	}
	for x := x1; x <= x2; x++ {
		oceanMap[y][x]++
	}
}

func drawVerticalLine(x, y1, y2 int) {
	if(y1 > y2) {
		y1, y2 = y2, y1
	}
	for y := y1; y <= y2; y++ {
		oceanMap[y][x]++
	}
}

func drawDiagonalLine(x1, y1, x2, y2 int) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)

	if(dx > dy) {
		if(x1 > x2) {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		for x := x1; x <= x2; x++ {
			y := y1 + (y2 - y1) * (x - x1) / dx
			oceanMap[y][x]++
		}
	} else {
		if(y1 > y2) {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		for y := y1; y <= y2; y++ {
			x := x1 + (x2 - x1) * (y - y1) / dy
			oceanMap[y][x]++
		}
	}
}

func cordsToInt(cords []string) (int, int) {
	x, _ := strconv.Atoi(strings.TrimSpace(cords[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(cords[1]))
	return x, y
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 5)
	if err != nil {
	log.Fatal(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
