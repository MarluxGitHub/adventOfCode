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
var year = 2019
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
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result += mass/3 - 2
	}
}

// Solve part 2
func Solve2() {
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		for mass > 0 {
			mass = mass/3 - 2
			if mass > 0 {
				result += mass
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
