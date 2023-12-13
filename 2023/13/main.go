package main

import (
	"MarluxGitHub/adventOfCode/pkg/math"
	luxStrings "MarluxGitHub/adventOfCode/pkg/strings"

	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 13

var patterns [][]string

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseInput()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	result = summarizeReflection("", false)
}

// Solve part 2
func Solve2() {
	result = summarizeReflection("", true)
}

func parseInput() {
	var blocks [][]string
	var block []string

	for _, line := range lines {
		if line == "" {
			blocks = append(blocks, block)
			block = nil
			continue
		}
		block = append(block, line)
	}
	blocks = append(blocks, block)

	patterns = blocks
}

func printPattern(pattern []string) {
	for _, row := range pattern {
		fmt.Println(row)
		fmt.Println("")
	}
	fmt.Println("")
}

func findReflection(grid []string, smudge bool) int {
	for rIdx := 1; rIdx < len(grid); rIdx++ {
		gridCopy := make([]string, len(grid))
		copy(gridCopy, grid)
		left := gridCopy[:rIdx]
		right := gridCopy[rIdx:]
		min := math.Min(len(left), len(right))

		left = left[len(left)-min:]
		right = right[:min]
		slices.Reverse(right)

		leftStr := strings.Join(left, "")
		rightStr := strings.Join(right, "")

		if (smudge && luxStrings.HammingDistance(leftStr, rightStr) == 1) || (!smudge && leftStr == rightStr) {
			return rIdx
		}
	}
	return -1
}

func summarizeReflection(inputData string, smudge bool) (total int) {
	for _, grid := range patterns {
		rowIdx := findReflection(grid, smudge)
		if rowIdx != -1 {
			total += rowIdx * 100
			continue
		}

		colIdx := findReflection(rotateGrid(grid), smudge)
		if colIdx != -1 {
			total += colIdx
			continue
		}
		panic("no reflection found")
	}
	return
}

// Rotates 90 degrees
func rotateGrid(grid []string) []string {
	newRowLength := len(grid[0])
	rotated := make([]string, newRowLength)
	for i := 0; i < newRowLength; i++ {
		for j := range grid {
			rotated[i] += string(grid[len(grid)-1-j][i])
		}
	}
	return rotated
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
