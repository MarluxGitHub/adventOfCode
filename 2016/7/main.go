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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2016
var day = 7

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

func Solve1() {
	for _, line := range lines {
		outer, inner := parseInput(line)

		innerABBA := false

		for _, i := range inner {
			if ContainsABBA(i) {
				innerABBA = true
				break
			}
		}

		if innerABBA {
			continue
		}

		for _, o := range outer {
			if ContainsABBA(o) {
				result++
				break
			}
		}
	}
}

func parseInput(line string) (outer, inner []string) {
	current := ""

	for _, rn := range line {
		switch char := string(rn); char {
		case "[":
			outer = append(outer, current)
			current = ""
		case "]":
			inner = append(inner, current)
			current = ""
		default:
			current += char
		}
	}

	outer = append(outer, current)

	return outer, inner
}

// Solve part 2
func Solve2() {

	for _, line := range lines {
		outer, inner := parseInput(line)

	outer:
		for _, o := range outer {
			pairs, ok := ContainsABA(o)

			if ok {
				for _, i := range inner {
					for _, pair := range pairs {
						if ContainsBAB(i, pair[0], pair[1]) {
							result++
							continue outer
						}
					}
				}
			}
		}
	}
}

func ContainsABBA(str string) bool {
	for i := 3; i < len(str); i++ {
		// match outsides, match insides, ensure insides and outsides are different
		if str[i-3] == str[i] && str[i-2] == str[i-1] && str[i] != str[i-1] {
			return true
		}
	}
	return false
}

func ContainsABA(str string) (pairs [][]string, ok bool) {
	ok = false
	for i := 2; i < len(str); i++ {
		// match outsides, match insides, ensure insides and outsides are different
		if str[i-2] == str[i] && str[i] != str[i-1] {
			pairs = append(pairs, []string{string(str[i]), string(str[i-1])})
			ok = true
		}
	}
	return pairs, ok
}

func ContainsBAB(str, a, b string) bool {
	for i := 2; i < len(str); i++ {
		// match outsides, match insides, ensure insides and outsides are different
		inner := str[i-2 : i+1]
		if inner == b+a+b {
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
