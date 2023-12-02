package main

import (
	"MarluxGitHub/adventOfCode/pkg/math"
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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 2


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
		gameRounds := strings.Split(line, ":")

		var gameNumber int
		// Parse Game x
		fmt.Sscanf(gameRounds[0], "Game %d", &gameNumber)

		rounds := strings.Split(gameRounds[1], ";")

		impossible := false

		out:
		for _, round := range rounds {
			roundPlays := strings.Split(round, ",")
			for _, roundPlay := range roundPlays {
				var number int = 0
				var color string = ""

				fmt.Sscanf(roundPlay, "%d %s", &number, &color)

				switch color {
				case "red":
					if number > 12 {
						impossible = true
					}
				case "green":
					if number > 13 {
						impossible = true
					}
				case "blue":
					if number > 14 {
						impossible = true
					}
				}

				if impossible {
					break out
				}
			}
		}

		if !impossible {
			result += gameNumber
		}
	}

}

// Solve part 2
func Solve2() {
	for _, line := range lines {
		gameRounds := strings.Split(line, ":")

		var gameNumber int
		// Parse Game x
		fmt.Sscanf(gameRounds[0], "Game %d", &gameNumber)

		rounds := strings.Split(gameRounds[1], ";")

		red := 0
		blue := 0
		green := 0

		for _, round := range rounds {
			roundPlays := strings.Split(round, ",")
			for _, roundPlay := range roundPlays {
				var number int = 0
				var color string = ""

				fmt.Sscanf(roundPlay, "%d %s", &number, &color)

				switch color {
				case "red":
					red = math.Max(red, number)
				case "green":
					green = math.Max(green,number)
				case "blue":
					blue = math.Max(blue, number)
				}
			}
		}

		result += red * green * blue
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
