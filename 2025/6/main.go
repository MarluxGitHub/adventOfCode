package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 6

type MathProblem struct {
	Numbers     []int
	Computation func(int, int) int
}

var MathProblems []MathProblem

var grid map[datastructures.Point]rune

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genGrid()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func genGrid() {
	grid = make(map[datastructures.Point]rune)

	for y, line := range lines {
		for x, r := range line {
			pt := datastructures.Point{X: x, Y: y}
			grid[pt] = r
		}
	}
}

func parseInput1() []MathProblem {
	MathProblems = nil

	// determine width/height
	width := 0
	for _, l := range lines {
		if len(l) > width {
			width = len(l)
		}
	}
	height := len(lines)

	// pad lines to same width to simplify indexing
	for i := range lines {
		if len(lines[i]) < width {
			lines[i] = lines[i] + strings.Repeat(" ", width-len(lines[i]))
		}
	}

	// helper to get rune at x,y (treat missing as space)
	runeAt := func(x, y int) rune {
		pt := datastructures.Point{X: x, Y: y}
		if r, ok := grid[pt]; ok {
			return r
		}
		return ' '
	}

	columnIsAllSpaces := func(x int) bool {
		for y := 0; y < height; y++ {
			if runeAt(x, y) != ' ' {
				return false
			}
		}
		return true
	}

	// scan columns and extract contiguous blocks (problems)
	for x := 0; x < width; {
		if columnIsAllSpaces(x) {
			x++
			continue
		}

		start := x
		for x < width && !columnIsAllSpaces(x) {
			x++
		}
		end := x - 1

		// operation is expected on the last row of the block
		opRow := height - 1
		var opRune rune = ' '
		for xx := start; xx <= end; xx++ {
			if r := runeAt(xx, opRow); r != ' ' {
				opRune = r
				break
			}
		}

		var comp func(int, int) int
		switch opRune {
		case '+':
			comp = func(a, b int) int { return a + b }
		case '*':
			comp = func(a, b int) int { return a * b }
		default:
			comp = nil
		}

		nums := []int{}
		// each row above opRow may contain a number (digits within the block)
		for y := 0; y < opRow; y++ {
			s := ""
			for xx := start; xx <= end; xx++ {
				r := runeAt(xx, y)
				if r >= '0' && r <= '9' {
					s += string(r)
				}
			}
			s = strings.TrimSpace(s)
			if s != "" {
				if n, err := strconv.Atoi(s); err == nil {
					nums = append(nums, n)
				}
			}
		}

		MathProblems = append(MathProblems, MathProblem{Numbers: nums, Computation: comp})
	}

	return MathProblems
}

func parseInput2() []MathProblem {
	// Cephalopod math: numbers are written right-to-left in columns.
	// Each number occupies a single column with most significant digit at the top.

	MathProblems = nil

	// determine width/height
	width := 0
	for _, l := range lines {
		if len(l) > width {
			width = len(l)
		}
	}
	height := len(lines)

	// pad lines to same width to simplify indexing
	for i := range lines {
		if len(lines[i]) < width {
			lines[i] = lines[i] + strings.Repeat(" ", width-len(lines[i]))
		}
	}

	runeAt := func(x, y int) rune {
		pt := datastructures.Point{X: x, Y: y}
		if r, ok := grid[pt]; ok {
			return r
		}
		return ' '
	}

	columnIsAllSpaces := func(x int) bool {
		for y := 0; y < height; y++ {
			if runeAt(x, y) != ' ' {
				return false
			}
		}
		return true
	}

	for x := 0; x < width; {
		if columnIsAllSpaces(x) {
			x++
			continue
		}

		start := x
		for x < width && !columnIsAllSpaces(x) {
			x++
		}
		end := x - 1

		opRow := height - 1
		var opRune rune = ' '
		for xx := start; xx <= end; xx++ {
			if r := runeAt(xx, opRow); r != ' ' {
				opRune = r
				break
			}
		}

		var comp func(int, int) int
		switch opRune {
		case '+':
			comp = func(a, b int) int { return a + b }
		case '*':
			comp = func(a, b int) int { return a * b }
		default:
			comp = nil
		}

		nums := []int{}
		// Read columns right-to-left; each column (top-to-bottom) is a number
		for xx := end; xx >= start; xx-- {
			s := ""
			for y := 0; y < opRow; y++ {
				r := runeAt(xx, y)
				if r >= '0' && r <= '9' {
					s += string(r)
				}
			}
			s = strings.TrimSpace(s)
			if s != "" {
				if n, err := strconv.Atoi(s); err == nil {
					nums = append(nums, n)
				}
			}
		}

		MathProblems = append(MathProblems, MathProblem{Numbers: nums, Computation: comp})
	}

	// mark this task done in our local state by returning the problems
	return MathProblems

}

// Solve part 1
func Solve1() {
	parseInput1()
	for _, mp := range MathProblems {
		if mp.Computation == nil || len(mp.Numbers) == 0 {
			continue
		}

		res := mp.Numbers[0]

		for i := 1; i < len(mp.Numbers); i++ {
			res = mp.Computation(res, mp.Numbers[i])
		}

		result += res
	}
}

// Solve part 2
func Solve2() {
	parseInput2()

	for _, mp := range MathProblems {
		if mp.Computation == nil || len(mp.Numbers) == 0 {
			continue
		}

		res := mp.Numbers[0]

		for i := 1; i < len(mp.Numbers); i++ {
			res = mp.Computation(res, mp.Numbers[i])
		}

		result += res
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
