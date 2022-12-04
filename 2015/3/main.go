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
var year = 2015
var day = 3


func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()

  readInput()

  solve1()
  println("1:" + strconv.Itoa(result))

  solve2()
  println("2:" + strconv.Itoa(result))
}

type point struct {
	x,y int
}


func solve1() {
	result =0

	x,y := 0,0
	houses := make(map[point]int)

	for _, line := range lines {
		for _, char := range line {
			switch char {
			case '^': y++
			case 'v': y--
			case '<': x--
			case '>': x++
			}

			houses[point{x,y}]++
		}
	}

	for _, v := range houses {
		if v > 0 {
			result++
		}
	}
}

func solve2() {
	result =0

	sx,sy,rx,ry := 0,0,0,0
	houses := make(map[point]int)

	for _, line := range lines {
		for i, char := range line {
			x,y := 0,0
			if i % 2 == 0 {
				x,y = sx, sy
			} else {
				x,y = rx, ry
			}

			switch char {
			case '^': y++
			case 'v': y--
			case '<': x--
			case '>': x++
			}

			if i % 2 == 0 {
				sx,sy = x, y
			} else {
				rx,ry = x, y
			}

			houses[point{x,y}]++
		}
	}

	result = len(houses)
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
