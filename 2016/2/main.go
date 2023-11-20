package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
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
var result string
var year = 2016
var day = 2

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	result = ""
	Solve1()
	println("1:" + result)

	result = ""
	Solve2()
	println("2:" + result)
}

// Solve part 1
func Solve1() {
	erg := 0
	keypad := buildKeypad()

	pos := datastructures.Point{0, 0}

	for _, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				if pos.Y < 1 {
					pos.Y++
				}
			case 'D':
				if pos.Y > -1 {
					pos.Y--
				}
			case 'R':
				if pos.X < 1 {
					pos.X++
				}
			case 'L':
				if pos.X > -1 {
					pos.X--
				}
			}
		}

		erg = erg*10 + keypad[pos]
	}
	result = strconv.Itoa(erg)
}

// Solve part 2
func Solve2() {
	result = ""

	keypad := buildKeyPad2()

	pos := datastructures.Point{-2, 0}

	for _, line := range lines {
		for _, c := range line {
			newpos := pos
			switch c {
			case 'U':
				newpos.Y++
			case 'D':
				newpos.Y--
			case 'R':
				newpos.X++
			case 'L':
				newpos.X--
			}

			if _, ok := keypad[newpos]; ok {
				pos = newpos
			}
		}

		result += string(keypad[pos])
	}
}

func buildKeypad() map[datastructures.Point]int {
	keypad := make(map[datastructures.Point]int)

	keypad[datastructures.Point{-1, 1}] = 1
	keypad[datastructures.Point{0, 1}] = 2
	keypad[datastructures.Point{1, 1}] = 3
	keypad[datastructures.Point{-1, 0}] = 4
	keypad[datastructures.Point{0, 0}] = 5
	keypad[datastructures.Point{1, 0}] = 6
	keypad[datastructures.Point{-1, -1}] = 7
	keypad[datastructures.Point{0, -1}] = 8
	keypad[datastructures.Point{1, -1}] = 9

	return keypad
}

//	  1
//	2 3 4
//
// 5 6 7 8 9
//
//	A B C
//	  D
func buildKeyPad2() map[datastructures.Point]rune {
	keypad := make(map[datastructures.Point]rune)

	keypad[datastructures.Point{0, 2}] = '1'
	keypad[datastructures.Point{-1, 1}] = '2'
	keypad[datastructures.Point{0, 1}] = '3'
	keypad[datastructures.Point{1, 1}] = '4'
	keypad[datastructures.Point{-2, 0}] = '5'
	keypad[datastructures.Point{-1, 0}] = '6'
	keypad[datastructures.Point{0, 0}] = '7'
	keypad[datastructures.Point{1, 0}] = '8'
	keypad[datastructures.Point{2, 0}] = '9'
	keypad[datastructures.Point{-1, -1}] = 'A'
	keypad[datastructures.Point{0, -1}] = 'B'
	keypad[datastructures.Point{1, -1}] = 'C'
	keypad[datastructures.Point{0, -2}] = 'D'

	return keypad
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
