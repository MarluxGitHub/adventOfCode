package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2024
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
		// search string for string mul(3 digits, 3 digits)
		regexp := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

		// find all matches

		matches := regexp.FindAllStringSubmatch(line, -1)

		// iterate over matches
		for _, match := range matches {
			// convert strings to integers
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])

			// add product to result
			result += a * b
		}
	}

}

// Solve part 2
func Solve2() {

	use := true
	for _, line := range lines {
		// search string for string mul(3 digits, 3 digits)
		regexp := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)

		// find all matches
		matches := regexp.FindAllStringSubmatch(line, -1)

		// iterate over matches
		for _, match := range matches {
			if match[1] == "" {
				if match[4] == "do()" {
					use = true
				} else if match[5] == "don't()" {
					use = false
				}
			} else if use {
				// convert strings to integers
				a, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])

				// add product to result
				result += a * b
			}
		}
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
