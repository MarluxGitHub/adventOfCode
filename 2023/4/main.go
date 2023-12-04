package main

import (
	luxMath "MarluxGitHub/adventOfCode/pkg/math"
	"bufio"
	"fmt"
	"log"
	"math"
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
var day = 4


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
	for _, line := range lines {
		winningNumbers, yourNumbers := parseLine(line)
		rightNumbers := luxMath.Intersection(winningNumbers, yourNumbers)

		if len(rightNumbers) == 0 {
			continue
		}

		points := int(math.Pow(2, float64(len(rightNumbers)-1)))

		result += points
	}
}

// Solve part 2
func Solve2() {

}

func parseLine(line string)(winningNumbers, yourNumbers []int){
	// throw away game x:
	line = strings.Split(line, ":")[1]

	numbers := strings.Split(line, "|")

	winningNumbers = parseNumbers(numbers[0])
	yourNumbers = parseNumbers(numbers[1])

	return winningNumbers, yourNumbers
}

func parseNumbers(numbers string) []int {
	stringNumbers := strings.Split(numbers, " ")
	res := make([]int, 0)

	for _, number := range stringNumbers {
		if number == "" {
			continue
		}

		n, err := strconv.Atoi(number)

		if err != nil {
			log.Fatal(err)
		}
		res = append(res, n)
	}

	return res
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
