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
var day = 4


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
		// read var from string in format elf1start-elf1finish,elf2start-elf2finish
		elf1start, elf1finish, elf2start, elf2finish := 0, 0, 0, 0
		fmt.Sscanf(line, "%d-%d,%d-%d", &elf1start, &elf1finish, &elf2start, &elf2finish)

		// get absolute start and finish of elf1 and elf2
		elf1Abs := elf1finish-elf1start
		elf2Abs := elf2finish-elf2start

		if elf2Abs > elf1Abs {
			// swap start and finish from elf1 and 2
			elf1start, elf1finish, elf2start, elf2finish = elf2start, elf2finish, elf1start, elf1finish
		}

		if(elf1start <= elf2start && elf1finish >= elf2finish) {
			result++
		}
	}
}

func solve2() {
	result = 0
	for _, line := range lines {
		// read var from string in format elf1start-elf1finish,elf2start-elf2finish
		elf1start, elf1finish, elf2start, elf2finish := 0, 0, 0, 0
		fmt.Sscanf(line, "%d-%d,%d-%d", &elf1start, &elf1finish, &elf2start, &elf2finish)

		// get absolute start and finish of elf1 and elf2
		elf1Abs := elf1finish-elf1start
		elf2Abs := elf2finish-elf2start

		if elf2Abs > elf1Abs {
			// swap start and finish from elf1 and 2
			elf1start, elf1finish, elf2start, elf2finish = elf2start, elf2finish, elf1start, elf1finish
		}

		// Check if elf1 start and finish and elf2 start and finish overlap
		if(elf1start <= elf2start && elf1finish >= elf2start) {
			result++
		} else if(elf1start <= elf2finish && elf1finish >= elf2finish) {
			result++
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
