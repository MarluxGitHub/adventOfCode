package main

import (
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
var year = 2023
var day = 5

type GardenConv struct {
	From, To string
	Converter func(int) int
}

var ConverterMap = map[string]GardenConv{}
var Seeds = []int{}


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()
	parseInput()

	result = 0
  	Solve1()
  	println("1:" + strconv.Itoa(result))

	result = 0
  	Solve2()
  	println("2:" + strconv.Itoa(result))
}

func parseInput() {
	parseSeeds()

	from, to := 0,0
	for i, line := range lines[2:] {
		if strings.Contains(line, "map:") {
			from = i
		}

		if strings.Trim(line, " ") == "" {
			to = i
		}

		if from != 0 && to != 0 {
			parseConverter(from,to)
			from,to = 0,0
		}
	}
}

func parseSeeds() {
	line := lines[0]

	seedsElements := strings.Split(strings.Split(line, ":")[1], " ")

	for _, seedElement := range seedsElements {
		seed, err := strconv.Atoi(seedElement)

		if err != nil {
			log.Fatal(err)
		}

		Seeds = append(Seeds, seed)
	}
}

func parseConverter(from, to int) {
	// x-to-y map:
	// parse x and y from line[from]

	var x, y string

	fmt.Sscanf(lines[from], "%s-to-%s map:", &x, &y)
}

// Solve part 1
func Solve1() {

}

// Solve part 2
func Solve2() {

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
