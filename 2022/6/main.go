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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2022
var day = 6


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

  	solve1()
  	println("1:" + strconv.Itoa(result))

  	solve2()
  	println("2:" + strconv.Itoa(result))
}

func solve1() {
	result = solve(4)
}

func solve2() {
  	result = solve(14)
  	// Solve part 2
}

func solve(indicator int) int {
	line := lines[0]

	for i := indicator; i < len(line); i++ {
		// Get a substring from i-indicator to i
		if hasOnlyUniqueChars(line[i-indicator : i]) {
			println(line[i-indicator : i])
			return i
		}
	}

	return 0
}

func hasOnlyUniqueChars(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
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
