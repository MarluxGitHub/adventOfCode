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
var year = 2023
var day = 15

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
		split := strings.Split(line, ",")
		for _, s := range split {
			if s != "" {
				result += HashCode(s)
			}
		}
	}
}

// Solve part 2
func Solve2() {
	boxes := make([][]string, 256)
	focalLengths := make(map[string]int)

	instructions := []string{}

	for _, line := range lines {
		split := strings.Split(line, ",")
		for _, s := range split {
			if s != "" {
				instructions = append(instructions, s)
			}
		}
	}

	for _, instruction := range instructions {
		if strings.Contains(instruction, "-") {
			label := instruction[:len(instruction)-1]
			index := HashCode(label)

			for i, l := range boxes[index] {
				if l == label {
					boxes[index] = append(boxes[index][:i], boxes[index][i+1:]...)
					break
				}
			}
		} else {
			parts := strings.Split(instruction, "=")
			label := parts[0]
			length := parts[1]

			lengthValue, err := strconv.Atoi(length)
			if err != nil {
				fmt.Println("Error parsing length:", err)
				return
			}

			index := HashCode(label)
			if !contains(boxes[index], label) {
				boxes[index] = append(boxes[index], label)
			}

			focalLengths[label] = lengthValue
		}
	}

	total := 0

	for boxNumber, box := range boxes {
		for lensSlot, label := range box {
			total += (boxNumber + 1) * (lensSlot + 1) * focalLengths[label]
		}
	}

	result = total
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func HashCode(s string) int {

	current := 0
	for _, r := range s {
		current += int(r)
		current *= 17
		current %= 256
	}

	return current
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
