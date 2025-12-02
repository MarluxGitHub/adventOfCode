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
var year = 2025
var day = 2

type IDRange struct {
	Left, Right int
}

var IDRanges []IDRange

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseIDRangesFromLine(lines[0])

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	for _, r := range IDRanges {
		for i := r.Left; i <= r.Right; i++ {
			if isNumberFake(i) {
				result += i
			}
		}
	}
}

// Solve part 2
func Solve2() {
	for _, r := range IDRanges {
		for i := r.Left; i <= r.Right; i++ {
			if isNumberFakeTwo(i) {
				result += i
			}
		}
	}
}

func parseIDRangesFromLine(line string) {
	for _, part := range strings.Split(line, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		IDRanges = append(IDRanges, parseIDRange(part))
	}
}

func parseIDRange(s string) IDRange {
	var r IDRange

	fmt.Sscanf(s, "%d-%d", &r.Left, &r.Right)

	return r
}

func isNumberFake(n int) bool {
	s := strconv.Itoa(n)

	// skip if uneven length
	if len(s)%2 != 0 {
		return false
	}

	// split string in left and right if uneven has the middle both parts
	// example 11 is left 1 and right 1
	mid := len(s) / 2
	left := s[:mid]
	right := s[len(s)-mid:]

	// if right has leading zero skip
	if strings.HasPrefix(right, "0") {
		return false
	}

	return left == right
}

func isNumberFakeTwo(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)

	// try every possible pattern length from 1 up to l/2
	for p := 1; p <= l/2; p++ {
		if l%p != 0 {
			continue
		}
		times := l / p
		pattern := s[:p]
		ok := true
		for i := 1; i < times; i++ {
			if s[i*p:(i+1)*p] != pattern {
				ok = false
				break
			}
		}
		if ok && times >= 2 {
			return true
		}
	}

	return false
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
