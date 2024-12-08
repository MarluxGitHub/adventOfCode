package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
var year = 2024
var day = 8

var antennaMap map[datastructures.Point]rune
var antennaDictionary map[rune][]datastructures.Point
var antiLocs map[datastructures.Point]bool

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
	generateAntennaMap()

	for _, cords := range antennaDictionary {
		if len(cords) == 1 {
			continue
		}

		// Check everycombination of cords
		for i := 0; i < len(cords); i++ {
			for j := i + 1; j < len(cords); j++ {
				// get distance between cords
				distance := cords[j].Subtract(cords[i])

				a1 := cords[i].Subtract(distance)
				a2 := cords[j].Add(distance)

				// check if a1 is in bounds
				if _, ok := antennaMap[a1]; ok {
					antiLocs[a1] = true
					if antennaMap[a1] == '.' {
						antennaMap[a1] = '#'
					}
				}

				// check if a2 is in bounds
				if _, ok := antennaMap[a2]; ok {
					antiLocs[a2] = true
					if antennaMap[a2] == '.' {
						antennaMap[a2] = '#'
					}
				}

			}
		}
	}

	for _, a := range antiLocs {
		if a {
			result++
		}
	}

}

// Solve part 2
func Solve2() {
	for _, cords := range antennaDictionary {
		// Check everycombination of cords
		for i := 0; i < len(cords); i++ {
			for j := i + 1; j < len(cords); j++ {
				// get distance between cords
				distance := cords[j].Subtract(cords[i])
				current := cords[i]

				for {
					current = current.Add(distance)

					if _, ok := antennaMap[current]; !ok {
						break
					}

					antiLocs[current] = true
					if antennaMap[current] == '.' {
						antennaMap[current] = '#'
					}

				}

				current = cords[j]

				for {
					current = current.Subtract(distance)

					if _, ok := antennaMap[current]; !ok {
						break
					}

					antiLocs[current] = true
					if antennaMap[current] == '.' {
						antennaMap[current] = '#'
					}

				}
			}
		}
	}

	for _, a := range antiLocs {
		if a {
			result++
		}
	}

}

func generateAntennaMap() {
	antennaMap = make(map[datastructures.Point]rune)
	antennaDictionary = make(map[rune][]datastructures.Point)
	antiLocs = make(map[datastructures.Point]bool)

	for y, line := range lines {
		for x, c := range line {
			antennaMap[datastructures.Point{X: x, Y: y}] = c

			if c != '.' {
				antennaDictionary[c] = append(antennaDictionary[c], datastructures.Point{X: x, Y: y})
			}
		}
	}
}

func printAntennaMap() {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			fmt.Print(string(antennaMap[datastructures.Point{X: x, Y: y}]))
		}
		fmt.Println()
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
