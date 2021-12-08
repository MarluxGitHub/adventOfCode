package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var input []ProblemSet

type ProblemSet struct {
	wires []string
	input []string
}

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()
  transformInput()
  result := solve()
  println(fmt.Sprintf("Result: %v", result))

}

func solve() int {
	sum := 0
	for _, ps := range input {
		sum += solveProblemSet(ps)
	}

	return sum
}

func solveProblemSet(ps ProblemSet) int {
	sum := 0

	for _, input := range ps.input {
		length := len(input)
		if(length == 2 || length == 3 || length == 4 || length == 7 ) {
			sum++
		}
	}

	return sum
}

func transformInput() {
	input = make([]ProblemSet, len(lines))
	for i, line := range lines {
		input[i] = parseLine(line)
	}
}

func parseLine(line string) ProblemSet {
	problem := strings.Split(line, " | ")

	wires := strings.Split(problem[0], " ")

	for i, wire := range wires {
		wires[i] = SortString(standardizeSpaces(wire))
	}

	inputs := strings.Split(problem[1], " ")

	for i, input := range inputs {
		inputs[i] = SortString(standardizeSpaces(input))
	}

	return ProblemSet{wires, inputs}
}

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func printProblemSet(ps ProblemSet) {
	println(fmt.Sprintf("wires: %v", ps.wires))
	println(fmt.Sprintf("input: %v", ps.input))
}


func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 8)
	if err != nil {
	log.Fatal(err)
	}
}
