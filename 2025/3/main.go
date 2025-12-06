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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 3

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
	max := 0

	for _, line := range lines {
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				battery := string(line[i]) + string(line[j])
				voltage, _ := strconv.Atoi(battery)
				if voltage > max {
					max = voltage
				}
			}
		}

		result += max
		max = 0
	}
}

// Solve part 2
func Solve2() {
	max := 0
	lasthit := 0
	currentVoltage := ""

	for _, line := range lines {
		newLasthit := -1
		for i := 12; i > 0; i-- {
			lasthit = newLasthit + 1
			for j := len(line) - i; j >= lasthit; j-- {
				battery := line[j]
				voltage, _ := strconv.Atoi(string(battery))
				if max <= voltage {
					max = voltage
					newLasthit = j
				}
			}
			currentVoltage += strconv.Itoa(max)
			max = 0
		}
		resultVoltage, _ := strconv.Atoi(currentVoltage)
		result += resultVoltage
		currentVoltage = ""
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
