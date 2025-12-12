package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 12

type Present struct {
	Shape map[datastructures.Point]bool
}

var Presents map[int]Present

type Space struct {
	Width, Height int
	Presents      map[int]int
}

var Spaces map[int]Space

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genPresentsAndSpaces()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func genPresentsAndSpaces() {
	Presents = make(map[int]Present)
	Spaces = make(map[int]Space)

	i := 0
	spaceIndex := 0

	// Parse Present-Shapes
	for i < len(lines) {
		line := lines[i]

		// Check if this is a present shape header (format: "INDEX:")
		if strings.HasSuffix(line, ":") {
			// Extract index
			indexStr := strings.TrimSuffix(line, ":")
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			i++ // Move to first shape line

			// Parse shape lines until empty line
			shape := make(map[datastructures.Point]bool)
			y := 0
			for i < len(lines) && lines[i] != "" {
				for x, char := range lines[i] {
					if char == '#' {
						shape[datastructures.Point{X: x, Y: y}] = true
					}
				}
				y++
				i++
			}

			Presents[index] = Present{Shape: shape}

			// Skip empty line
			if i < len(lines) && lines[i] == "" {
				i++
			}
		} else if strings.Contains(line, "x") {
			// This is a space line (format: "WIDTHxHEIGHT: q0 q1 q2 q3 q4 q5")
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				log.Fatal("Invalid space format:", line)
			}

			// Parse dimensions
			dimParts := strings.Split(parts[0], "x")
			if len(dimParts) != 2 {
				log.Fatal("Invalid dimension format:", parts[0])
			}
			width, err := strconv.Atoi(dimParts[0])
			if err != nil {
				log.Fatal(err)
			}
			height, err := strconv.Atoi(dimParts[1])
			if err != nil {
				log.Fatal(err)
			}

			// Parse present quantities
			quantities := strings.Fields(parts[1])
			presents := make(map[int]int)
			for idx, qtyStr := range quantities {
				qty, err := strconv.Atoi(qtyStr)
				if err != nil {
					log.Fatal(err)
				}
				if qty > 0 {
					presents[idx] = qty
				}
			}

			Spaces[spaceIndex] = Space{
				Width:    width,
				Height:   height,
				Presents: presents,
			}
			spaceIndex++
			i++
		} else {
			i++
		}
	}
}

// Solve part 1
func Solve1() {
	// Berechne die Flächen der Presents (shapes)
	shapes := make([]int, len(Presents))
	for idx, present := range Presents {
		shapes[idx] = len(present.Shape)
	}

	// Für jeden Space prüfen
	for _, space := range Spaces {
		// Fläche des Spaces berechnen
		area := space.Width * space.Height

		// Gesamtfläche der Presents berechnen
		total := 0
		for presentIdx, quant := range space.Presents {
			if presentIdx < len(shapes) {
				total += quant * shapes[presentIdx]
			}
		}

		// Wenn total < area, dann zählen
		if total < area {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
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
