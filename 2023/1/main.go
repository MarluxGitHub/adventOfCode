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
var year = 2023
var day = 1

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
	solve(false)
}

// Solve part 2
func Solve2() {
	solve(true)
}

func solve(two bool) {
	for _, line := range lines {
		if two {
			line = replaceNumberWords(line)
		}
		number := ""
		for _, c := range line {
			// check if c is 0 - 9
			if c >= 48 && c <= 57 {
				number += string(c)
			}
		}

		if number != "" {
			// get only the first and last rune
			number = string(number[0]) + string(number[len(number)-1])

			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			result += n
		}
	}
}

func replaceNumberWords(s string) string {
	// map of number words to their number
	// e.g. "one" -> "o1e"
	// o and e -> for overlapping words
	numberWords := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for k, v := range numberWords {
		s = strings.ReplaceAll(s, k, v)
	}

	return s
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
