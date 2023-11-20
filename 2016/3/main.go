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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2016
var day = 3

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
	for _, line := range lines {
		a, b, c := 0, 0, 0
		fmt.Sscanf(line, "%d %d %d", &a, &b, &c)

		result += isPossible(a, b, c)
	}
}

// Solve part 2
func Solve2() {
	for i := 0; i < len(lines); i += 3 {
		a1, b1, c1 := 0, 0, 0
		a2, b2, c2 := 0, 0, 0
		a3, b3, c3 := 0, 0, 0

		fmt.Sscanf(lines[i], "%d %d %d", &a1, &b1, &c1)
		fmt.Sscanf(lines[i+1], "%d %d %d", &a2, &b2, &c2)
		fmt.Sscanf(lines[i+2], "%d %d %d", &a3, &b3, &c3)

		result += isPossible(a1, a2, a3)
		result += isPossible(b1, b2, b3)
		result += isPossible(c1, c2, c3)
	}
}

func isPossible(a, b, c int) int {
	if a+b > c && a+c > b && b+c > a {
		return 1
	}

	return 0
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
