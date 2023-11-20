package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
var year = 2022
var day = 10

var crt [][]rune

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
	x := 1
	cpuCycles := 0

	stack := input2Fifo()

	for !stack.IsEmpty() {
		cpuCycles++
		current := cpuCycles - 20

		if current%40 == 0 {
			result += cpuCycles * x
		}

		val := stack.Pop()
		x += val.(int)
	}
}

// Solve part 2
func solve2() {
	stack := input2Fifo()
	x := 1
	cpuCycles := 0

	// make rune array 40x6 filled with '.'

	crt = make([][]rune, 6)
	for i := range crt {
		crt[i] = make([]rune, 40)
	}

	for i := range crt {
		for j := range crt[i] {
			crt[i][j] = ' '
		}
	}

	for !stack.IsEmpty() {
		row := cpuCycles / 40
		col := cpuCycles % 40

		if cpuCycles%40 >= x-1 && cpuCycles%40 <= x+1 {
			crt[row][col] = '#'
		}

		val := stack.Pop()
		x += val.(int)

		cpuCycles++
	}

	// print crt
	for i := range crt {
		for j := range crt[i] {
			printf(string(crt[i][j]))
		}
		println("")
	}
}

func input2Fifo() *datastructures.Fifo {
	fifo := datastructures.NewFIFO()

	for _, line := range lines {
		args := strings.Split(line, " ")
		fifo.Push(0)

		if args[0] == "addx" {
			val, err := strconv.Atoi(args[1])

			if err != nil {
				log.Fatal(err)
			}

			fifo.Push(val)
		}
	}

	return fifo
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
