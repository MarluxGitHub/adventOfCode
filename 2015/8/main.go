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
var year = 2015
var day = 8

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
	result = 0

	sumOfChars := 0
	sumOfMemory := 0

	for _, line := range lines {
		sumOfChars += len(line)

		l := ""

		for i := 1; i < len(line)-1; i++ {
			if line[i] == '\\' {
				if line[i+1] == 'x' {
					l += " "
					i += 3
				} else {
					l += " "
					i++
				}
			} else {
				l += string(line[i])
			}
		}

		sumOfMemory += len(l)
	}

	result = sumOfChars - sumOfMemory
}

// Solve part 2
func Solve2() {
	result = 0

	sumOfChars := 0
	sumOfMemory := 0

	for _, line := range lines {
		sumOfChars += len(line)

		for i := 0; i < len(line); i++ {
			if line[i] == '\\' || line[i] == '"' {
				sumOfMemory++
			}

			sumOfMemory++
		}

		sumOfMemory += 2
	}

	result = sumOfMemory - sumOfChars
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
