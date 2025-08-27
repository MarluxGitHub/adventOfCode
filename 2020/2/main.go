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
var year = 2020
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
	for _, i := range lines {
		min := 0
		max := 0
		letter := 'a'
		password := ""

		// scan line für min-max letter: password
		fmt.Sscanf(i, "%d-%d %c: %s", &min, &max, &letter, &password)

		count := 0
		for _, c := range password {
			if c == letter {
				count++
			}
		}

		if count >= min && count <= max {
			result++
		}
	}
}

// Solve part 2
func Solve2() {
	for _, i := range lines {
		first, second := 0, 0
		password := ""
		b := 'a'

		fmt.Sscanf(i, "%d-%d %c: %s", &first, &second, &b, &password)
		letter := byte(b)

		firstLetter := password[first-1]
		secondLetter := password[second-1]

		if ((firstLetter == letter) ||
			(secondLetter == letter)) &&
			firstLetter != secondLetter {
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
