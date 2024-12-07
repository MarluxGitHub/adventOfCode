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
var result string
var year = 2024
var day = 7

type equation struct {
	erg     int
	numbers []int
}

var equations []equation

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseInput()

	result = ""
	Solve1()
	println("1:" + result)

	result = ""
	Solve2()
	println("2:" + result)
}

// Solve part 1
func Solve1() {
	total, err := totalEquationCalibration([]func(*int, int) int{Add, Multiply})
	if err != nil {
		log.Fatal(err)
	}
	result = strconv.Itoa(total)
}

// Solve part 2
func Solve2() {
	total, err := totalEquationCalibration([]func(*int, int) int{Add, Multiply, Concatenate})
	if err != nil {
		log.Fatal(err)
	}
	result = strconv.Itoa(total)
}

// Add operator: sums the accumulator and term.
func Add(acc *int, term int) int {
	if acc == nil {
		return term
	}
	return *acc + term
}

// Multiply operator: multiplies the accumulator and term.
func Multiply(acc *int, term int) int {
	if acc == nil {
		return term
	}
	return *acc * term
}

// Concatenate operator: appends the current term to the accumulator.
func Concatenate(acc *int, term int) int {
	if acc == nil {
		return term
	}
	accStr := strconv.Itoa(*acc)
	termStr := strconv.Itoa(term)
	concatenated, _ := strconv.Atoi(accStr + termStr)
	return concatenated
}

// ValidateEquation recursively validates if the result can be achieved using the given terms and operators.
func ValidateEquation(result int, terms []int, acc *int, operators []func(*int, int) int) bool {
	// Base case
	if len(terms) == 0 {
		return acc != nil && *acc == result
	}

	// Recursive case
	for _, op := range operators {
		newAcc := op(acc, terms[0])
		if ValidateEquation(result, terms[1:], &newAcc, operators) {
			return true
		}
	}
	return false
}

// Part1 solves part 1 using Add and Multiply operators.
func totalEquationCalibration(funcs []func(*int, int) int) (int, error) {
	total := 0
	for _, eq := range equations {
		result := eq.erg
		terms := eq.numbers
		if ValidateEquation(result, terms, nil, funcs) {
			total += result
		}
	}
	return total, nil
}

func parseInput() {
	for _, line := range lines {
		// line is erg: num1 num2 num3 ... numx

		splitted := strings.Split(line, " ")
		// remove the last symbol from splitted 0

		erg, _ := strconv.Atoi(splitted[0][:len(splitted[0])-1])

		numbers := make([]int, len(splitted)-1)

		for i := 1; i < len(splitted); i++ {
			numbers[i-1], _ = strconv.Atoi(splitted[i])
		}

		equations = append(equations, equation{erg, numbers})

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
