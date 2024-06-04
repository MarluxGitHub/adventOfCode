package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"

	intcode "MarluxGitHub/adventOfCode/pkg/intCode"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2019
var day = 2

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
	strIntCodes := strings.Split(lines[0], ",")
	intCodes := make([]int, len(strIntCodes))
	for i, strIntCode := range strIntCodes {
		intCodes[i], _ = strconv.Atoi(strIntCode)
	}

	intCodes[1] = 12
	intCodes[2] = 2

	intCodeProcessor := intcode.NewInterpreter(intCodes)
	intCodeProcessor.Run()

	result = intCodeProcessor.GetValueOfRegister(0)
}

// Solve part 2
func Solve2() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			strIntCodes := strings.Split(lines[0], ",")
			intCodes := make([]int, len(strIntCodes))
			for i, strIntCode := range strIntCodes {
				intCodes[i], _ = strconv.Atoi(strIntCode)
			}

			intCodes[1] = noun
			intCodes[2] = verb

			intCodeProcessor := intcode.NewInterpreter(intCodes)
			intCodeProcessor.Run()

			if intCodeProcessor.GetValueOfRegister(0) == 19690720 {
				result = 100*noun + verb
				return
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
