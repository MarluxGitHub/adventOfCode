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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 3

var engineMap = map[datastructures.Point]*EnginePart{}

type EnginePart struct {
	Id int
	IsNumber,IsNumberStart bool
	IsRune bool
	IsGear bool
	Value int
}

func newEnginePart() *EnginePart {
	return &EnginePart{0,false,false, false, false,0}
}

func initInput() {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			engineMap[datastructures.Point{i,j}] = newEnginePart()
		}
	}
}

func parseInput() {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				engineMap[datastructures.Point{i,j}].IsNumber = true
				if engineMap[datastructures.Point{i,j-1}] == nil || !engineMap[datastructures.Point{i,j-1}].IsNumber {
					engineMap[datastructures.Point{i,j}].IsNumberStart = true
				}
			} else if lines[i][j] == '.' {
				continue
			} else if lines[i][j] == '*' {
				engineMap[datastructures.Point{i,j}].IsGear = true
				engineMap[datastructures.Point{i,j}].IsRune = true
			} else {
				engineMap[datastructures.Point{i,j}].IsRune = true
			}
		}
	}

	id := 1
	for key, value := range engineMap {
		if value.IsNumberStart {
			val := 0
			for i := key.Y; i < len(lines[key.X]); i++ {
				if engineMap[datastructures.Point{key.X,i}].IsNumber {
					val = val * 10 + int(lines[key.X][i] - '0')
				} else {
					break
				}
			}

			for i := key.Y; i < len(lines[key.X]); i++ {
				if engineMap[datastructures.Point{key.X,i}].IsNumber {
					engineMap[datastructures.Point{key.X,i}].Value = val
					engineMap[datastructures.Point{key.X,i}].Id = id
				} else {
					break
				}
			}
		}
		id++
	}
}

func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()
	initInput()
	parseInput()

	result = 0
  	Solve1()
  	println("1:" + strconv.Itoa(result))

	result = 0
  	Solve2()
  	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	out:
	for key, value := range engineMap {
		if value.IsNumberStart {
			for i := key.Y; engineMap[datastructures.Point{key.X,i}] != nil && engineMap[datastructures.Point{key.X,i}].IsNumber == true; i++ {
				for j := key.X - 1; j <= key.X + 1; j++ {
					for k := i - 1; k <= i + 1; k++ {
						if engineMap[datastructures.Point{j,k}] == nil {
							continue
						}
						if engineMap[datastructures.Point{j,k}].IsRune {
							result += value.Value
							continue out
						}
					}
				}
			}
		}
	}
}

// Solve part 2
func Solve2() {
	for key, value := range engineMap {
		if value.IsGear {
			partMap := map[int]*EnginePart{}

			// Manhattan Matrix Check for PartIds
			for i := key.X - 1; i <= key.X + 1; i++ {
				for j := key.Y - 1; j <= key.Y + 1; j++ {
					if engineMap[datastructures.Point{i,j}] == nil {
						continue
					}
					if engineMap[datastructures.Point{i,j}].IsNumber {
						partMap[engineMap[datastructures.Point{i,j}].Id] = engineMap[datastructures.Point{i,j}]
					}
				}
			}

			if len(partMap) == 2 {
				res := 1
				for _, value := range partMap {
					res *= value.Value
				}

				result += res
			}
		}
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
