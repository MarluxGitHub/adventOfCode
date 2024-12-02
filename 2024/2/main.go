package main

import (
	"MarluxGitHub/adventOfCode/pkg/math"
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

	for _, line := range lines {
		// its a space separated string of numbers so we split it
		strnums := strings.Split(line, " ")

		// convert strnums to ints
		nums := make([]int, len(strnums))
		for i, strnum := range strnums {
			nums[i], _ = strconv.Atoi(strnum)
		}

		if solver(nums, false) {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
	for _, line := range lines {
		// its a space separated string of numbers so we split it
		stringNumbers := strings.Split(line, " ")

		// convert stringNumbers to integers
		numbers := make([]int, len(stringNumbers))
		for i, stringNumber := range stringNumbers {
			numbers[i], _ = strconv.Atoi(stringNumber)
		}

		if solver(numbers, true) {
			result++
		}
	}
}

func solver(nums []int, isFirst bool) bool {
	bit := math.Sign(nums[1] - nums[0])

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if bit != math.Sign(diff) || math.Abs(diff) <= 0 || math.Abs(diff) >= 4 {
			if isFirst {
				// generate new number slices with each time a number is removed
				for j := 0; j < len(nums); j++ {
					nums2 := make([]int, len(nums)-1)
					copy(nums2, nums[:j])
					copy(nums2[j:], nums[j+1:])
					if solver(nums2, false) {
						return true
					}
				}
			}

			return false
		}
	}

	return true
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
