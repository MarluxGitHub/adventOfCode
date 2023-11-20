package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"

	"MarluxGitHub/adventOfCode/pkg/datastructures"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2022
var day = 9

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	result = 0
	solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	solve2()
	println("2:" + strconv.Itoa(result))
}

// Advent of Code 2022 Day 9 Part 1
func solve1() {
	result = solveRopeOfLenN(2)
}

func solve2() {
	result = solveRopeOfLenN(10)
}

func solveRopeOfLenN(n int) int {
	rope := make([]datastructures.Point, n)

	for i := 0; i < n; i++ {
		rope[i] = datastructures.Point{X: 0, Y: 0}
	}

	visited := make(map[datastructures.Point]bool)

	visited[rope[n-1]] = true

	for _, line := range lines {
		args := strings.Split(line, " ")
		direction := args[0]
		distance, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				rope[0].Y++
			case "D":
				rope[0].Y--
			case "R":
				rope[0].X++
			case "L":
				rope[0].X--
			}

			for j := 1; j < n; j++ {
				if rope[j-1].MooreDistance(rope[j]) > 1 {
					rope[j] = rope[j].Add(rope[j-1].Subtract(rope[j]).Normalize())
				}
			}

			visited[rope[n-1]] = true
		}
	}

	return len(visited)
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
