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
var input []int

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()
  transformInput()
  response := minCostToMakeElementEqual(input, len(input))
  println(strconv.Itoa(response))
}

func transformInput() {
	i := strings.Split(lines[0], ",")
	input = make([]int, len(i))
	for k, _ := range i {
		input[k], _ = strconv.Atoi(i[k])
	}
}

func computeCost(arr []int, n int, x int) int {
	cost := 0
	for i := 0; i < n; i++ {
		steps := abs(arr[i] - x)
		cost += sumFunction(steps)
	}
	return cost
}

func sumFunction(n int) int {
	return n * (n + 1) / 2
}

func minCostToMakeElementEqual(a []int, n int) int {
	minCost := computeCost(a, n, a[0])
	for i := 1; i < n; i++ {
		cost := computeCost(a, n, i)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 7)
	if err != nil {
	log.Fatal(err)
	}
}
