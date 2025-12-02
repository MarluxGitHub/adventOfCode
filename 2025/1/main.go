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
var year = 2025
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
	currentPoint := 50

	for _, line := range lines {
		direction := line[:1]
		s := line[1:]

		steps, _ := strconv.Atoi(s)

		if direction == "L" {
			steps = -steps
		}

		currentPoint = (currentPoint + steps + 100) % 100

		if currentPoint == 0 {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
	currentPoint := 50

	for _, line := range lines {
		direction := line[:1]
		s := line[1:]

		steps, _ := strconv.Atoi(s)

		for i := 0; i < steps; i++ {
			if direction == "L" {
				currentPoint--
			} else {
				currentPoint++
			}

			currentPoint = (currentPoint + 100) % 100

			if currentPoint == 0 {
				result++
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
