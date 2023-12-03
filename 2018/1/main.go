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
var year = 2018
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
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		result += i
	}
}

// Solve part 2
func Solve2() {
	mapSignals := make(map[int]bool)
	sum := 0

	for {
		for _, line := range lines {
			i, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			sum += i
			if mapSignals[sum] {
				result = sum
				return
			}
			mapSignals[sum] = true
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
