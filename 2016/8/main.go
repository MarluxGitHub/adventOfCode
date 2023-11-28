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
var year = 2016
var day = 8

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	result = 0
	Solve()
	println("1:" + strconv.Itoa(result))
}

var display map[datastructures.Point]bool

// Solve part 1
func Solve() {
	initDisplay()

	for _, line := range lines {
		parseLine(line)
	}

	for _, v := range display {
		if v {
			result++
		}
	}

	printDisplay()
}

func initDisplay() {
	display = make(map[datastructures.Point]bool)

	for i := 0; i < 50; i++ {
		for j := 0; j < 6; j++ {
			display[datastructures.Point{i, j}] = false
		}
	}
}

func parseLine(line string) {
	operation := strings.Split(line, " ")[0]

	switch operation {
	case "rect":
		rect(line)
	case "rotate":
		rotate(line)
	default:
		break
	}
}

func rect(line string) {
	var x, y int
	fmt.Sscanf(line, "rect %dx%d", &x, &y)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			display[datastructures.Point{i, j}] = true
		}
	}
}

func rotate(line string) {

	var s, x string
	var y int
	var amount int

	fmt.Sscanf(line, "rotate %s %s by %d", &s, &x, &amount)
	y, _ = strconv.Atoi(x[2:])

	switch s {
	case "row":
		rotateRow(y, amount)
	case "column":
		rotateColumn(y, amount)
	default:
		break
	}
}

func rotateRow(y, amount int) {
	for i := 0; i < amount; i++ {
		rotateRowOnce(y)
	}
}

func rotateRowOnce(y int) {
	var last bool
	for i := 0; i < 50; i++ {
		if i == 0 {
			last = display[datastructures.Point{i, y}]
			display[datastructures.Point{i, y}] = display[datastructures.Point{49, y}]
		} else {
			current := display[datastructures.Point{i, y}]
			display[datastructures.Point{i, y}] = last
			last = current
		}
	}
}

func rotateColumn(x, amount int) {
	for i := 0; i < amount; i++ {
		rotateColumnOnce(x)
	}
}

func rotateColumnOnce(x int) {
	var last bool
	for i := 0; i < 6; i++ {
		if i == 0 {
			last = display[datastructures.Point{x, i}]
			display[datastructures.Point{x, i}] = display[datastructures.Point{x, 5}]
		} else {
			current := display[datastructures.Point{x, i}]
			display[datastructures.Point{x, i}] = last
			last = current
		}
	}
}

func printDisplay() {
	for j := 0; j < 6; j++ {
		for i := 0; i < 50; i++ {
			if display[datastructures.Point{i, j}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
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
