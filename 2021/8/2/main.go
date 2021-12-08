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

	tempLen := make(map[int]string)
	problems := mergeTwoArrayUnique(ps.wires, ps.input)

	tempLen[1] = findLen(problems, 2)[0]
	tempLen[4] = findLen(problems, 4)[0]
	tempLen[7] = findLen(problems, 3)[0]
	tempLen[8] = findLen(problems, 7)[0]

	// 9
	for _, s := range findLen(problems, 6) {
		if len(stringDiff(s, tempLen[4])) == 2 {
			tempLen[9] = s
		}
	}

	// 6
	for _, s := range findLen(problems, 6) {
		if len(stringDiff(s, tempLen[7])) == 4 {
			tempLen[6] = s
		}
	}

	// 5
	for _, s := range findLen(problems, 5) {
		if len(stringDiff(tempLen[6], s)) == 1 {
			tempLen[5] = s
		}
	}

	// 3
	for _, s := range findLen(problems, 5) {
		if len(stringDiff(tempLen[9], s)) == 1 && s != tempLen[5] {
			tempLen[3] = s
		}
	}

	for _, s := range findLen(problems, 6) {
		if len(stringDiff(tempLen[8], s)) == 1 && s != tempLen[9] && s != tempLen[6] {
			tempLen[0] = s
		}
	}

	for _, s := range problems {
		hit := false
		for _, v := range tempLen {
			if s == v {
				hit = true
			}
		}
		if !hit {
			tempLen[2] = s
		}
	}

	for _, v := range ps.input {
		for z, s := range tempLen {
			if(s == v) {
				sum = sum * 10 + z
			}
		}
	}

	return sum
}

func findLen(inputs []string, l int) []string {
	res := []string{}
	for _, i := range inputs {
		if len(i) == l {
			res = append(res, i)
		}
	}
	return res
}

func stringDiff(s1, s2 string) string {
	result := ""
	for _, c1 := range s1 {
		in := false
		for _, c2 := range s2 {
			if c1 == c2 {
				in = true
			}
		}
		if !in {
			result += string(c1)
		}
	}
	return result
}

func mergeTwoArrayUnique(a1, a2 []string) []string {
	m := make(map[string]bool)
	for _, v := range a1 {
		m[v] = true
	}
	for _, v := range a2 {
		m[v] = true
	}

	var merged []string
	for k := range m {
		merged = append(merged, k)
	}

	return merged
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
