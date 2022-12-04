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
	for _, line := range lines {
		for _, c := range line {
			if c == '(' {
				result++
			} else {
				result--
			}
		}
	}
}

func solve2() {
	result = 0

	for _, line := range lines {
		for i, c := range line {
			if c == '(' {
				result++
			} else {
				result--
			}

			if result == -1 {
				result = i+1
				return
			}
		}
		break
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
