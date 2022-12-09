package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"

	"MarluxGitHub/adventOfCode/internal/datastructures"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

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
	head := datastructures.Point{X: 0, Y: 0}
	tail := datastructures.Point{X: 0, Y: 0}

	visited := make(map[datastructures.Point]bool)

	visited[tail] = true

	for _, line := range lines {
		args := strings.Split(line, " ")
		direction := args[0]
		distance, err := strconv.Atoi(args[1])

		if(err != nil) {
			log.Fatal(err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "R":
				head.X++
			case "L":
				head.X--
			}

			if head.MooreDistance(tail) > 1 {
				tail = tail.Add(head.Subtract(tail).Normalize())
			}

			visited[tail] = true
		}
	}

	result = len(visited)
}

func solve2() {
	rope := make([]datastructures.Point, 10)

	for i := 0; i < 10; i++ {
		rope[i] = datastructures.Point{X: 0, Y: 0}
	}

	visited := make(map[datastructures.Point]bool)

	visited[rope[9]] = true

	for _, line := range lines {
		args := strings.Split(line, " ")
		direction := args[0]
		distance, err := strconv.Atoi(args[1])

		if(err != nil) {
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

			for j := 1; j < 10; j++ {
				if rope[j-1].MooreDistance(rope[j]) > 1 {
					rope[j] = rope[j].Add(rope[j-1].Subtract(rope[j]).Normalize())
				}
			}


			visited[rope[9]] = true
		}
	}

	result = len(visited)
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
