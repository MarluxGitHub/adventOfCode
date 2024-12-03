package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/echojc/aocutil" // external package for Advent of Code utilities
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }

// printf is not used, so it has been removed

var lines []string
var result int
var year = 2015
var day = 11

var next string

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
	next = getNextPassword(lines[0])

	for !passWordRuleValidator(next) {
		next = getNextPassword(next)
	}

	println(next)
}

// Solve part 2
func Solve2() {
	next = getNextPassword(next)

	for !passWordRuleValidator(next) {
		next = getNextPassword(next)
	}

	println(next)
}

func getNextPassword(password string) string {
	bytes := []byte(password)
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] == 'z' {
			bytes[i] = 'a'
		} else {
			bytes[i]++
			break
		}
	}
	return string(bytes)
}

func passWordRuleValidator(password string) bool {
	return passwordRuleOne(password) && passwordRuleTwo(password) && passwordRuleThree(password)
}

// must include one increasing straight of at least three letters
func passwordRuleOne(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1]-1 && password[i] == password[i+2]-2 {
			return true
		}
	}

	return false
}

// don't contain i, o, l
// dont contain i,o,l
func passwordRuleTwo(password string) bool {
	for _, c := range password {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
	}

	return true
}

// must contain at least two different, non-overlapping pairs of letters
func passwordRuleThree(password string) bool {
	pairs := 0
	i := 0
	for i < len(password)-1 {
		if password[i] == password[i+1] {
			pairs++
			i += 2
		} else {
			i++
		}
	}

	return pairs >= 2
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
