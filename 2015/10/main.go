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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2015
var day = 10

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
	res := lines[0]

	for i := 0; i < 40; i++ {
		res = solver(res)
	}

	result = len(res)
}

// Solve part 2
func Solve2() {
	res := lines[0]

	for i := 0; i < 50; i++ {
		res = solver(res)
	}

	result = len(res)
}

func solver(line string) string {
	if len(line) == 0 {
		return ""
	}

	current := rune(line[0])
	count := 0

	var result strings.Builder

	for i := 0; i < len(line); i++ {
		if current == rune(line[i]) {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteRune(current)
			current = rune(line[i])
			count = 1
		}
	}

	if count > 0 {
		result.WriteString(strconv.Itoa(count))
		result.WriteRune(current)
	}

	return result.String()
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
