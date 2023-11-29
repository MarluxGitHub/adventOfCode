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
var year = 2016
var day = 9

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
	compressed := getCompressedLine()
	result = decode(compressed, false)
}

// Solve part 2
func Solve2() {
	compressed := getCompressedLine()
	result = decode(compressed, true)
}

func decode(compressed string, solveSubstring bool) int {

	count := 0

	for i := 0; i < len(compressed); i++ {
		if compressed[i] == '(' {
			// Get the marker
			marker := ""
			for compressed[i] != ')' {
				marker += string(compressed[i])
				i++
			}
			marker += string(compressed[i])

			// Get the marker values
			markerValues := getMarkerValues(marker)

			// Get the string to repeat
			stringToRepeat := ""
			for j := 0; j < markerValues[0]; j++ {
				i++
				stringToRepeat += string(compressed[i])
			}

			if solveSubstring {
				count += decode(stringToRepeat, true) * markerValues[1]
			} else {
				count += len(stringToRepeat) * markerValues[1]
			}

		} else {
			count++
		}
	}

	return count
}

func getMarkerValues(marker string) []int {
	markerValues := []int{0, 0}
	marker = marker[1 : len(marker)-1]

	for i := 0; i < len(marker); i++ {
		if marker[i] == 'x' {
			markerValues[0], _ = strconv.Atoi(marker[:i])
			markerValues[1], _ = strconv.Atoi(marker[i+1:])
			break
		}
	}

	return markerValues
}

func getCompressedLine() string {
	compressed := ""

	for _, line := range lines {
		compressed += line
	}

	return compressed
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
