package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2015
var day = 12

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
	result = getSumOfAllNumbersInString(lines[0])
}

// Solve part 2
func Solve2() {
	reg := "\"red\""
	regex := regexp.MustCompile(reg)

	for {
		i := regex.FindStringIndex(lines[0])

		if i == nil {
			break
		}

		stack := datastructures.NewStack()
		left := -1
		right := -1

		isCurlyBracket := false

	left:
		for j := i[0] - 1; j >= 0; j-- {
			switch lines[0][j] {
			case '{':
				if stack.IsEmpty() {
					left = j
					isCurlyBracket = true
					break left
				} else {
					stack.Pop()
				}
			case '}':
				stack.Push('}')
			case '[':
				if stack.IsEmpty() {
					left = j
					break left
				} else {
					stack.Pop()
				}
			case ']':
				stack.Push(']')
			}
		}

	right:
		for j := i[1]; j < len(lines[0]); j++ {
			switch lines[0][j] {
			case '{':
				stack.Push('}')
			case '}':
				if stack.IsEmpty() {
					isCurlyBracket = true
					right = j
					break right
				}
				stack.Pop()
			case '[':
				stack.Push(']')
			case ']':
				if stack.IsEmpty() {
					right = j
					break right
				}
				stack.Pop()
			}
		}

		if left == -1 || right == -1 {
			continue
		}

		if isCurlyBracket {
			lines[0] = lines[0][:left] + lines[0][right+1:]
		} else {
			substring := lines[0][left : right+1]
			sum := getSumOfAllNumbersInString(substring)
			lines[0] = lines[0][:left] + "[" + strconv.Itoa(sum) + "]" + lines[0][right+1:]
		}
	}

	// print lines[0] in test.json
	os.WriteFile("test.json", []byte(lines[0]), 0644)

	result = getSumOfAllNumbersInString(lines[0])
}

func getSumOfAllNumbersInString(s string) int {
	// get All Numbers in the String
	reg := "(-?\\d+)"
	regexp := regexp.MustCompile(reg)

	// get all the matches
	matches := regexp.FindAllString(s, -1)

	sum := 0
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		sum += num
	}

	return sum
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
