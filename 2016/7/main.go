package main

import (
	"MarluxGitHub/adventOfCode/pkg/strings"
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

// Solve part 1
func Solve1() {
	for _,line := range lines {
		a,b,c := getStringsOfLine(line)

		if (!strings.ContainsABBA(b) && (strings.ContainsABBA(a) || strings.ContainsABBA(c))) {
			result++
		}
	}
}

// Solve part 2
func Solve2() {

}

func getStringsOfLine(line string) (left,inner,right string) {
	left = ""
	inner = ""
	right = ""
	innerMode := false
	rightMode := false

	for _,r := range line {
		if r == '[' {
			innerMode = true
			continue
		}

		if r == ']' {
			innerMode = false
			rightMode = true
			continue
		}

		if innerMode {
			inner += string(r)
		} else if rightMode {
			right += string(r)
		} else {
			left += string(r)
		}
	}

	return left,inner,right
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
