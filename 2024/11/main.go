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
var year = 2024
var day = 11

var numbers []int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genStones()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {

	for _, n := range numbers {
		result += int(traverse(uint64(n), 25))
	}
}

// Solve part 2
func Solve2() {
	for _, n := range numbers {
		result += int(traverse(uint64(n), 75))
	}
}

var memo = make(map[[2]int]uint64)

func traverse(stone uint64, steps int) uint64 {
	if steps == 0 {
		return 1
	}
	if val, ok := memo[[2]int{int(stone), steps}]; ok {
		return val
	}

	str := strconv.FormatUint(stone, 10)
	var result uint64

	if stone == 0 {
		result = traverse(1, steps-1)
	} else if len(str)%2 == 0 {
		left, _ := strconv.ParseUint(str[:len(str)/2], 10, 64)
		right, _ := strconv.ParseUint(str[len(str)/2:], 10, 64)
		result = traverse(left, steps-1) + traverse(right, steps-1)
	} else {
		result = traverse(stone*2024, steps-1)
	}

	memo[[2]int{int(stone), steps}] = result
	return result
}

func genStones() {
	numbersStr := strings.Split(lines[0], " ")

	numbers = make([]int, len(numbersStr))

	for i, n := range numbersStr {
		numbers[i], _ = strconv.Atoi(n)
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
