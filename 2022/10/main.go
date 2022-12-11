package main

import (
	"MarluxGitHub/adventOfCode/internal/datastructures"
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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2022
var day = 10


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

	result = 0
  	solve1()
  	println("1:" + strconv.Itoa(result))

	result = 0
  	solve2()
  	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func solve1() {
	x := 0
	cpuCycles := 0

	stack := datastructures.NewFIFO()

	for _, line := range lines {
		args := strings.Split(line, " ")
		stack.Push(0)

		if args[0] == "addx"  {
			val, err := strconv.Atoi(args[1])

			if err != nil {
				log.Fatal(err)
			}

			stack.Push(val)
		}
	}

	for !stack.IsEmpty() {
		cpuCycles++
		current := cpuCycles - 20

		if current % 40 == 0 {
			println("x:" + strconv.Itoa(x) + " cpuCycles:" + strconv.Itoa(cpuCycles))
			result += cpuCycles * x
		}

		val := stack.Pop()
		println("val:" + strconv.Itoa(val.(int)))
		x += val.(int)
	}
}

// Solve part 2
func solve2() {

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
