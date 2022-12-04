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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2015
var day = 5


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

  	solve1()
  	println("1:" + strconv.Itoa(result))

  	solve2()
  	println("2:" + strconv.Itoa(result))
}

func solve1() {
	result = 0

	for _, line := range lines {
		if isNiceString1(line) {
			result++
		}
	}
}

func solve2() {
  	result = 0

	for _, line := range lines {
		if isNiceString2(line) {
			result++
		}
	}
}

func isNiceString1(s string) bool {
	vowelcount := 0
	doubleletter := false

	for i, char := range s {
		switch char {
		case 'e', 'i', 'o', 'u':
			vowelcount++
		case 'a':
			vowelcount++
			if i < len(s)-1 && s[i+1] == 'b' {
				return false
			}
		case 'c':
			if i < len(s)-1 && s[i+1] == 'd' {
				return false
			}
		case 'p':
			if i < len(s)-1 && s[i+1] == 'q' {
				return false
			}
		case 'x':
			if i < len(s)-1 && s[i+1] == 'y' {
				return false
			}
		}
		if i > 0 && s[i-1] == s[i] {
			doubleletter = true
		}
	}

	return vowelcount >= 3 && doubleletter
}

func isNiceString2(s string) bool {
	pair := false
	repeat := false

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			repeat = true
		}
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i + 2; j < len(s)-1; j++ {
			if s[i] == s[j] && s[i+1] == s[j+1] {
				pair = true
			}
		}
	}

	return pair && repeat
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
