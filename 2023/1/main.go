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
var year = 2023
var day = 1


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
		number := ""
		for _, c := range line {
			// check if c is 0 - 9
			if c >= 48 && c <= 57 {
				number += string(c)
			}
		}

		if number != "" {
			// get only the first and last rune
			number = string(number[0]) + string(number[len(number)-1])

			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			result += n
		}
	}
}

// Solve part 2
func Solve2() {
	for _, line := range lines {
		line = replaceNumberWords(line)
		number := ""
		for _, c := range line {
			// check if c is 0 - 9
			if c >= 48 && c <= 57 {
				number += string(c)
			}
		}

		if number != "" {
			// get only the first and last rune
			number = string(number[0]) + string(number[len(number)-1])
			println(number)

			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			result += n
		}
	}
}

func replaceNumberWords(s string) string {
    numberWords := map[string]string{
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

	  return s
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
