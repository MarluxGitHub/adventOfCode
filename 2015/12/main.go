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
	// get All Numbers in the String
	for c := 0; c < len(lines[0]); c++ {
		if lines[0][c] == '-' || (lines[0][c] >= '0' && lines[0][c] <= '9') {
			num := ""
			for ; c < len(lines[0]) && (lines[0][c] == '-' || (lines[0][c] >= '0' && lines[0][c] <= '9')); c++ {
				num += string(lines[0][c])
			}
			i, _ := strconv.Atoi(num)
			result += i
		}
	}
}

// Solve part 2
func Solve2() {
	// find all positions of the word "red" in the string
	for c := 0; c < len(lines[0]); c++ {
		if strings.HasPrefix(lines[0][c:], `"red"`) {
			// go left until you hit a { or [
			l := c
			for ; l >= 0 && lines[0][l] != '{' && lines[0][l] != '['; l-- {
			}

			if lines[0][l] == '{' {
				// go right until you hit a }
				r := c + 4
				for ; r < len(lines[0]) && lines[0][r] != '}'; r++ {
				}

				if r < len(lines[0]) {
					lines[0] = lines[0][:l] + lines[0][r+1:]
					c = l
				}
			}
		}
	}

	for c := 0; c < len(lines[0]); c++ {
		if lines[0][c] == '-' || (lines[0][c] >= '0' && lines[0][c] <= '9') {
			num := ""
			for ; c < len(lines[0]) && (lines[0][c] == '-' || (lines[0][c] >= '0' && lines[0][c] <= '9')); c++ {
				num += string(lines[0][c])
			}
			i, _ := strconv.Atoi(num)
			result += i
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
