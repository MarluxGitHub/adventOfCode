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
var year = 2022
var day = 3

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()

  readInput()

  solve1()
  println("1:" + strconv.Itoa(result))

  result = 0

  solve2()
  println("2:" + strconv.Itoa(result))
}

func solve1() {
	for _, line := range lines {
		backpack1 := line[0:len(line)/2]
		backpack2 := line[len(line)/2:]

		// Find the first rune which is also in backpack2
		out:
		for _, rune1 := range backpack1 {
			for _, rune2 := range backpack2 {
				if rune1 == rune2 {
					value := int(rune1)

					if (value >= 65 && value <= 90) {value -= 38}
					if (value >= 97 && value <= 122) {value -= 96}

					result += value
					break out

				}
			}
		}
	}
}

func solve2() {
	for i := 0; i < len(lines); i += 3 {
		backpack1 := lines[i]
		backpack2 := lines[i+1]
		backpack3 := lines[i+2]

		// Find the first rune which is also in backpack2
		out:
		for _, rune1 := range backpack1 {
			for _, rune2 := range backpack2 {
				for _, rune3 := range backpack3 {
					if rune1 == rune2 && rune2 == rune3 {
						value := int(rune1)

						if (value >= 65 && value <= 90) {value -= 38}
						if (value >= 97 && value <= 122) {value -= 96}

						result += value
						break out
					}
				}
			}
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
