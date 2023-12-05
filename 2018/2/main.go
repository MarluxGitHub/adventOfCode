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

var lines []string
var year = 2018
var day = 2

var correctBoxes = make([]string, 0)

func main() {
	defer writer.Flush()

	readInput()

	var result = 0

	Solve1(&result)
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2(&result)

	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1(result *int) {
	twos := 0
	threes := 0

	for _, line := range lines {
		mapRune := make(map[rune]int)
		for _, r := range line {
			mapRune[r]++
		}

		two, three := false, false

		for _, v := range mapRune {
			if v == 2 {
				two = true
			}

			if v == 3 {
				three = true
			}
		}

		if two || three {
			correctBoxes = append(correctBoxes, line)
		}

		if two {
			twos++
		}

		if three {
			threes++
		}
	}

	*result = twos * threes
}

// Solve part 2
func Solve2(result *int) {
	for _, line := range correctBoxes {
		for _, line2 := range correctBoxes {
			if line == line2 {
				continue
			}

			diff := 0
			for i, r := range line {
				if r != rune(line2[i]) {
					diff++
				}
			}

			if diff == 1 {
				println(line)
				println(line2)
				*result = 0
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
