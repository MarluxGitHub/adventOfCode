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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 12

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	for _, line := range lines {
		springs := strings.Split(line, " ")[0]
		groupsStr := strings.Split(line, " ")[1]
		groups := []int{}
		for _, g := range strings.Split(groupsStr, ",") {
			number, _ := strconv.Atoi(g)
			groups = append(groups, number)
		}

		result += f(springs, groups, []int{0, 0, 0}, map[string]int{})
	}
}

// Solve part 2
func Solve2() {
	for _, line := range lines {
		springs := strings.Split(line, " ")[0]
		groupsStr := strings.Split(line, " ")[1]
		groups := []int{}
		for _, g := range strings.Split(groupsStr, ",") {
			number, _ := strconv.Atoi(g)
			groups = append(groups, number)
		}

		springsarray := make([]string, 5)
		for i := 0; i < 5; i++ {
			springsarray[i] = springs
		}

		springs = strings.Join(springsarray, "?")

		// repeat groups 5 times
		newGroups := make([]int, len(groups)*5)
		for i := 0; i < 5; i++ {
			copy(newGroups[i*len(groups):], groups)
		}
		groups = newGroups

		result += f(springs, groups, []int{0, 0, 0}, map[string]int{})
	}
}

func f(S string, G []int, state []int, cache map[string]int) int {
	stateKey := strings.Join(strings.Fields(fmt.Sprint(state)), ":")

	if val, ok := cache[stateKey]; ok {
		return val
	}

	_s := state[0]
	_g := state[1]
	length := state[2]

	if _s == len(S) {
		if _g == len(G)-1 && length == G[_g] {
			_g++
			length = 0
		}
		if _g == len(G) && length == 0 {
			return 1
		}
		return 0
	}

	result := 0

	if strings.Contains(".?", string(S[_s])) {
		if length == 0 {
			result += f(S, G, []int{_s + 1, _g, 0}, cache)
		} else if _g < len(G) && G[_g] == length {
			result += f(S, G, []int{_s + 1, _g + 1, 0}, cache)
		}
	}

	if strings.Contains("#?", string(S[_s])) {
		result += f(S, G, []int{_s + 1, _g, length + 1}, cache)
	}

	cache[stateKey] = result
	return result
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
