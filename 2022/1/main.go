package main

import (
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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2022
var day = 1


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
	current := 0

	for _, line := range lines {
		if line == "" {
			result = max(result, current)
			current = 0
		} else {
			calories, _ := strconv.Atoi(line)
			current += calories

		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve2() {
  	result = 0
	current := 0
	topElves := []int{}

	for _, line := range lines {
		if line == "" {
			topElves = append(topElves, current)
			current = 0
		} else {
			calories, _ := strconv.Atoi(line)
			current += calories
		}
	  }
	  sort.Slice(topElves, func(i, j int) bool {
		return topElves[i] > topElves[j]
	  })

	  // first 3 elements of result
	  if(len(topElves) > 3) {
		topElves = topElves[:3]
	  }

	  for _, v := range topElves {
	  	result += v
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
