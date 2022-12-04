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
var year = 2015
var day = 2


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
	for _, line := range lines {
		l,w,h := 0,0,0
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		result += computePaper(l,w,h)
	}
}

func solve2() {
	result = 0
	for _, line := range lines {
		l,w,h := 0,0,0
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		result += computeRibbon(l,w,h)
	}
}

func computePaper(l,w,h int) int {
  return 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
}

func computeRibbon(l,w,h int) int {
  order := []int{l,w,h}
  sort.Ints(order)
  return 2*order[0] + 2*order[1] + l*w*h
}

func min(nums ...int) int {
  min := nums[0]
  for _, n := range nums {
	if n < min {
	  min = n
	}
  }
  return min
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
