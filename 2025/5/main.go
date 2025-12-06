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
var day = 5

type IDRange struct {
	Left, Right int
}

var IDRanges []IDRange
var Ingredients []int

func readRange(line string) IDRange {
	IdRange := IDRange{}

	fmt.Sscanf(line, "%d-%d", &IdRange.Left, &IdRange.Right)

	return IdRange
}

func parseInput() {
	var ranges []IDRange
	br := 0

	for i, line := range lines {
		if line == "" {
			br = i + 1
			break
		}
		ranges = append(ranges, readRange(line))
	}

	IDRanges = mergeRanges(ranges)

	for i := br; i < len(lines); i++ {
		var ingredient int
		fmt.Sscanf(lines[i], "%d", &ingredient)
		Ingredients = append(Ingredients, ingredient)
	}
}

func mergeRanges(ranges []IDRange) []IDRange {
	if len(ranges) == 0 {
		return ranges
	}

	// merge all overlaps as long as some exist
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(ranges); i++ {
			for j := i + 1; j < len(ranges); j++ {
				// Check if ranges[i] and ranges[j] overlap
				if ranges[i].Left <= ranges[j].Right && ranges[i].Right >= ranges[j].Left {
					// Merge using outer bounds
					if ranges[j].Left < ranges[i].Left {
						ranges[i].Left = ranges[j].Left
					}
					if ranges[j].Right > ranges[i].Right {
						ranges[i].Right = ranges[j].Right
					}
					// Remove ranges[j]
					ranges = append(ranges[:j], ranges[j+1:]...)
					merged = true
					break
				}
			}
			if merged {
				break
			}
		}
	}

	return ranges
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseInput()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	for _, ingredient := range Ingredients {
		valid := false

		for _, r := range IDRanges {
			if ingredient >= r.Left && ingredient <= r.Right {
				valid = true
				break
			}
		}

		if valid {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
	for _, r := range IDRanges {
		result += r.Right - r.Left + 1
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
