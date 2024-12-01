package main

import (
	"MarluxGitHub/adventOfCode/pkg/math"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2024
var day = 1

// Make 2 Slices l1, l2 of type int
var l1, l2 []int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	for _, line := range lines {
		// number whitespace number
		n1, n2 := 0, 0
		fmt.Sscanf(line, "%d %d", &n1, &n2)
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	// copy l1 and l2 to local variables

	list1 := make([]int, len(l1))
	list2 := make([]int, len(l2))

	copy(list1, l1)
	copy(list2, l2)

	// sort list1 and list2

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		result += math.Abs(list1[i] - list2[i])
	}
}

// Solve part 2
func Solve2() {
	// make a map of int to int
	m := make(map[int]int)

	for i := 0; i < len(l2); i++ {
		m[l2[i]]++
	}

	for i := 0; i < len(l1); i++ {
		result += l1[i] * m[l1[i]]
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
