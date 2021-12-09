package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var heightMap[][]int


func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()
  linesToHeightMap()
  printHeightMap()
  result := solve()

  println("")

  println(fmt.Sprintf("Result: %v", result))
}

func solve() int {
	sum := 0

	sum += iterateHeightMapMaskSize2()

	return sum
}

func iterateHeightMapMaskSize2() int {
	sum := 0

	println("")
	println("")

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			hit := true
			out:
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					if i+x < 0 || i+x >= len(heightMap) {
						continue
					}
					if j+y < 0 || j+y >= len(heightMap[0]) {
						continue
					}

					if(heightMap[i+x][j+y] <= heightMap[i][j]) {
						hit = false
						break out;
					}
				}
			}

			if hit {
				println(fmt.Sprintf("%v,%v : %v", i, j, heightMap[i][j]))
				sum += heightMap[i][j] + 1
			}
		}
	}

	return sum
}

func linesToHeightMap() {
	heightMap = make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, c := range line {
			heightMap[i][j] = int(c - '0')
		}
	}
}

func printHeightMap() {
	for _, row := range heightMap {
		println("")
		for _, col := range row {
			printf(fmt.Sprintf("%v ", col))
		}
	}
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 9)
	if err != nil {
	log.Fatal(err)
	}
}
