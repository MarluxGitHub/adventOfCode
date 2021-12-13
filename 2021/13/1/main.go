package main

import (
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
var result int = 0

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()

  points, folds := parseLines(lines)
  results := len(folder(points, folds[0]))

  println(strconv.Itoa(results))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 13)
	if err != nil {
	log.Fatal(err)
	}
}


func folder(points []point, fold point) map[point]int {
	countByPoint := make(map[point]int)
	if fold.x == 0 {
		// Fold at y
		for _, p := range points {
			if p.y > fold.y {
				p.y = 2*fold.y - p.y

			}
			countByPoint[p]++
		}
	} else {
		// Fold at x
		for _, p := range points {
			if p.x > fold.x {
				p.x = 2*fold.x - p.x
			}
			countByPoint[p]++
		}
	}
	return countByPoint
}



type point struct {
	x, y int
}

func parseLines(lines []string) ([]point, []point) {
	isFoldInstructions := false
	var folds []point
	var points []point
	for _, line := range lines {
		if line == "" {
			isFoldInstructions = true
			continue
		}
		if isFoldInstructions {
			fold := strings.TrimSpace(strings.ReplaceAll(line, "fold along ", ""))
			parts := strings.Split(fold, "=")
			dir, val := parts[0], toInt(parts[1])
			var p point
			switch dir {
			case "x":
				p.x = val
			case "y":
				p.y = val
			}
			folds = append(folds, p)
		} else {
			parts := strings.Split(line, ",")
			x, y := toInt(parts[0]), toInt(parts[1])
			points = append(points, point{x: x, y: y})
		}
	}
	return points, folds
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}