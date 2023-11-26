package main

import (
	"MarluxGitHub/adventOfCode/pkg/strings"
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result string
var year = 2016
var day = 6


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

	result = ""
  	Solve1()
  	println("1:" + result)

	result = ""
  	Solve2()
  	println("2:" + result)
}

// Solve part 1
func Solve1() {
	columns := make([]string, len(lines[0]))

	for _,line := range lines {
		for i,r := range line {
			columns[i] += string(r)
		}
	}

	for _,column := range columns {
		result += string(strings.GetMostCommonRune(column))
	}
}

// Solve part 2
func Solve2() {
	columns := make([]string, len(lines[0]))

	for _,line := range lines {
		for i,r := range line {
			columns[i] += string(r)
		}
	}

	for _,column := range columns {
		result += string(strings.GetLeastCommonRune(column))
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
