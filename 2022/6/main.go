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
var year = 2022
var day = 6


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
	// is only 1 line but hey :)
    out:
	for _, line := range lines {
		for i := 4; i < len(line); i++ {
			// Get a substring from i-4 to i
			unique := hasOnlyUniqueChars(line[i-4 : i])
			if unique {
				println(line[i-4 : i])
				result = i
				break out
			}
		}
	}
}

func hasOnlyUniqueChars(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func solve2() {
  	result = 0
  	// Solve part 2
	out:
	for _, line := range lines {
	  for i := 14; i < len(line); i++ {
		  // Get a substring from i-4 to i
		  unique := hasOnlyUniqueChars(line[i-14 : i])
		  if unique {
			  println(line[i-14 : i])
			  result = i
			  break out
		  }
	  }
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
