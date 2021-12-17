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
var result int = 0

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()
	readInput()

	result = solve(217,240,-126,-69)

	println(strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err = i.Strings(2021, 17)
	if err != nil {
		log.Fatal(err)
	}
}

func solve(minX, maxX, minY, maxY int) int {
	uniq := make(map[point]bool)
	dy := minY
	if dy < 0 {
		dy = -dy
	}
	dy *= 10

	for x := 1; x <= maxX; x++ {
		for y := -dy; y < dy; y++ {
			vel := point{x: x, y: y}

			p := point{}
			for p.x+vel.x <= maxX && p.y+vel.y >= minY {
				p.x += vel.x
				p.y += vel.y
				vel.y -= 1
				if vel.x < 0 {
					vel.x += 1
				} else if vel.x > 0 {
					vel.x -= 1
				}
			}
			if p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY {
				uniq[point{x: x, y: y}] = true
			}
		}
	}

	return len(uniq)
}


type point struct {
	x, y int
}