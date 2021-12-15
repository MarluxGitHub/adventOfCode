package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }

var lines []string
var result = 0

var steps = 0
var octopusCount = 0

var octopuses [][]int
var flashMap [][]int

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()

  stringLinesTo2dIntArray()

  for {
		steps++
		if(step() == octopusCount) {
			break
		}
  }

  result = steps

  println("Result: " + strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 11)
	if err != nil {
	log.Fatal(err)
	}
}

func stringLinesTo2dIntArray () {
	octopuses = make([][]int, len(lines))
	flashMap = make([][]int, len(lines))
	for i, line := range lines {
		octopuses[i] = make([]int, len(line))
		flashMap[i] = make([]int, len(line))
		for j, char := range line {
			octopuses[i][j] = int(char) - 48
			flashMap[i][j] = 0
			octopusCount++
		}
	}
}

func step() int {
	flashOctopuses()

	flashes := countFlashMapValuesBiggerZero()
	clearFlashMap()

	return flashes
}

func flashOctopuses () {
	for i, row := range octopuses {
		for j := range row {
			flashOctopus(i, j)
		}
	}
}

func flashOctopus(x int, y int) {
	if(x < 0 || x >= len(octopuses)) {
		return
	}
	if(y < 0 || y >= len(octopuses[x])) {
		return
	}

	octopuses[x][y] += 1

	if(octopuses[x][y] > 9 && flashMap[x][y] == 0) {
		flashMap[x][y] = 1

		for xi := x - 1; xi <= x + 1; xi++ {
			for yi := y - 1; yi <= y + 1; yi++ {
				flashOctopus(xi, yi)
			}
		}
	}
}

func countFlashMapValuesBiggerZero()int {
	count := 0
	for i, row := range flashMap {
		for j := range row {
			if(flashMap[i][j] > 0) {
				count++
			}
		}
	}
	return count
}

func clearFlashMap () {
	for i := range flashMap {
		for j := range flashMap[i] {
			flashMap[i][j] = 0
			if(octopuses[i][j] > 9) {
				octopuses[i][j] = 0
			}
		}
	}
}
