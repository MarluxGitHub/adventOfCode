package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }

var lines []string
var result int = 0

var problem [][] int


func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()
	readInput()
	stringLinesTo2dIntArray()
	result = solve()

	println(strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 15)
	if err != nil {
	log.Fatal(err)
	}
}

func stringLinesTo2dIntArray () {
	problem = make([][]int, len(lines))

	for i, line := range lines {
		problem[i] = make([]int, len(line))

		for j, char := range line {
			problem[i][j] = int(char) - 48
		}
	}
}

func solve() int {
	graph := dijkstra.NewGraph()

	for i, row := range problem {
		for j := range row {
			graph.AddVertex(i * len(row) + j)
		}
	}

	for i, row := range problem {
		for j, value := range row {
			index := i * len(row) + j
			if checkLegitPoint(i-1, j) {
				graph.AddArc((i-1) * len(row) + j, index,  int64(value))
			}
			if checkLegitPoint(i+1, j) {
				graph.AddArc((i+1) * len(row) + j, index,  int64(value))
			}
			if checkLegitPoint(i, j-1) {
				graph.AddArc((i) * len(row) + j-1, index,  int64(value))
			}
			if checkLegitPoint(i, j+1) {
				graph.AddArc((i) * len(row) + j+1, index,  int64(value))
			}
		}
	}

	best, err := graph.Shortest(0, len(problem) * len(problem[0]) - 1)
	if err!=nil{
		log.Fatal(err)
	}

	return int(best.Distance)
}

func checkLegitPoint(x, y int) bool {
	if x < 0 || x >= len(problem) {
		return false
	}

	if y < 0 || y >= len(problem[x]) {
		return false
	}

	return true
}
