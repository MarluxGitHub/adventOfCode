package main

import (
	"MarluxGitHub/adventOfCode/internal/binary"
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
var year = 2015
var day = 7

var wires = make(map[string]int)


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
		var left, operand, right, destination string
		arguments := strings.Split(line, " -> ")

		left = arguments[0]
		destination = arguments[1]

		arguments = strings.Split(left, " ")

		if len(arguments) == 1 {
			left = arguments[0]
			operand = "VALUE"
		}

		if len(arguments) == 2 {
			left = arguments[1]
			operand = arguments[0]
		}

		if len(arguments) == 3 {
			left = arguments[0]
			operand = arguments[1]
			right = arguments[2]
		}

		leftValue := getValueOfWire(left)
		rightValue := getValueOfWire(right)

		println(line)
		println("leftvalue: " + strconv.Itoa(leftValue) + ", rightValue:" + strconv.Itoa(rightValue) + ", operand:" + operand)
		switch operand {
			case "VALUE":
				wires[destination] = leftValue
			case "NOT":
				wires[destination] = binary.SixteenBitNot(leftValue)
			case "AND":
				wires[destination] = binary.SixteenBitAnd(leftValue,rightValue)
			case "OR":
				wires[destination] = binary.SixteenBitOr(leftValue,rightValue)
			case "LSHIFT":
				wires[destination] = binary.SixteenBitLeftShift(leftValue,rightValue)
			case "RSHIFT":
				wires[destination] = binary.SixteenBitRightShift(leftValue,rightValue)
		}
		println(destination + " " + strconv.Itoa(wires[destination]))
		println("")

	}

	result = wires["a"]
}

func solve2() {
  	result = 0
  	// Solve part 2
}

func getValueOfWire(wire string) int {
	value, err := strconv.Atoi(wire)
	if err == nil {
		return value
	}

	return wires[wire]
}


func printwires() {
	for key, value := range wires {
		println(key + " " + strconv.Itoa(value))
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
