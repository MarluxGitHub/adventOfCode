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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2024
var day = 5

var rules map[int][]int
var indexPages int
var correctPages []int

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

// Solve part 1
func Solve1() {
	correct := 0

	for i := indexPages; i < len(lines); i++ {
		isCorrect := true

		numbersStr := strings.Split(lines[i], ",")

		//convert to int
		numbers := make([]int, len(numbersStr))
		for i, number := range numbersStr {
			numbers[i], _ = strconv.Atoi(number)
		}

		for _, number := range numbers {
			index1 := getIndexOfNumberInSlice(numbers, number)
			rulesOfIndex := rules[number]

			for _, rule := range rulesOfIndex {
				index2 := getIndexOfNumberInSlice(numbers, rule)
				if index1 > -1 && index2 > -1 && index1 > index2 {
					isCorrect = false
					break
				}
			}
		}

		if isCorrect {
			correct++
			mid := len(numbers) / 2
			result += numbers[mid]
			correctPages = append(correctPages, i)
		}
	}
}

func getIndexOfNumberInSlice(slice []int, number int) int {
	for i, value := range slice {
		if value == number {
			return i
		}
	}

	return -1
}

// Solve part 2
func Solve2() {
	// remove Correct pages
	for i := len(correctPages) - 1; i >= 0; i-- {
		lines = append(lines[:correctPages[i]], lines[correctPages[i]+1:]...)
	}

	for i := indexPages; i < len(lines); i++ {
		numbersStr := strings.Split(lines[i], ",")

		//convert to int
		numbers := make([]int, len(numbersStr))
		for j, number := range numbersStr {
			numbers[j], _ = strconv.Atoi(number)
		}

		swapped := true
		for swapped {
			swapped = false
			for j, number := range numbers {
				rulesOfIndex := rules[number]
				for _, rule := range rulesOfIndex {
					index1 := j
					index2 := getIndexOfNumberInSlice(numbers, rule)
					if index2 > -1 && index1 > index2 {
						// swap the numbers
						numbers[index1], numbers[index2] = numbers[index2], numbers[index1]
						swapped = true
					}
				}
			}
		}

		mid := len(numbers) / 2
		result += numbers[mid]
		correctPages = append(correctPages, i)
	}
}

func parseInput() {
	rules = make(map[int][]int)

	for i, line := range lines {
		if line == "" {
			indexPages = i + 1
			break
		}
		num1, num2 := 0, 0

		fmt.Sscanf(line, "%d|%d", &num1, &num2)

		// check if num1 is already in the map
		if _, ok := rules[num1]; ok {
			rules[num1] = append(rules[num1], num2)
		} else {
			rules[num1] = []int{num2}
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
