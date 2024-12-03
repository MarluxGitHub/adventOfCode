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
var year = 2015
var day = 12

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
	// get All Numbers in the String
	reg := "(-?\\d+)"
	regexp := regexp.MustCompile(reg)

	// get all the matches
	matches := regexp.FindAllString(lines[0], -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		result += num
	}
}

// Solve part 2
func Solve2() {
	// get all Locations of the String Red

	for {
		// find the first occurence of the word red
		redIndex := -1
		for i, c := range lines[0] {
			if c == '"' && i < len(lines[0])-2 && lines[0][i+1] == 'r' && lines[0][i+2] == 'e' && lines[0][i+3] == 'd' && lines[0][i+4] == '"' {
				redIndex = i
				break
			}
		}

		// if no red is found, break
		if redIndex == -1 {
			break
		}

		// find the first { and }
		openIndex := -1
		closeIndex := -1
		bracket := false
		for i := redIndex - 5; i >= 0; i-- {
			if lines[0][i] == '{' {
				openIndex = i
				break
			}

			if lines[0][i] == ']' {
				bracket = true
			}

			if lines[0][i] == '[' {
				if bracket {
					bracket = false
				} else {
					break
				}
			}
		}

		bracket = false
		for i := redIndex; i < len(lines[0]); i++ {
			if lines[0][i] == '}' {
				closeIndex = i
				break
			}

			if lines[0][i] == '[' {
				bracket = true
			}

			if lines[0][i] == ']' {
				if bracket {
					bracket = false
				} else {
					break
				}
			}
		}

		// remove the string from the openIndex to the closeIndex
		if openIndex != -1 && closeIndex != -1 {
			lines[0] = lines[0][:openIndex] + lines[0][closeIndex+1:]
		} else {
			break
		}
	}

	// get All Numbers in the String
	reg := "(-?\\d+)"
	regex := regexp.MustCompile(reg)

	// get all the matches
	matches := regex.FindAllString(lines[0], -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		result += num
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
