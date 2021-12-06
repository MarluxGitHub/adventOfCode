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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var fish [9]int

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()

  initFish()

  for i := 0; i < 256; i++ {
	simulateDay()
  }

  println(strconv.Itoa(countFish()))
}

func countFish() int {
	count := 0
	for i := 0; i < 9; i++ {
		count += fish[i]
	}
	return count
}

func printFish() {
	for i := 0; i < 9; i++ {
		printf(strconv.Itoa(fish[i]))
	}
	println("")
}

func initFish() {
	input := strings.Split(lines[0], ",")

	// for input
	for i := 0; i < len(input); i++ {
		key, _ := strconv.Atoi(input[i])
		fish[key]++
	}
}

func simulateDay() {
	fish2 := [9]int{}
	help := fish[0]

	for i := 1; i < 9; i++ {
		fish2[i-1] = fish[i]
	}

	fish = fish2
	fish[8] = help
	fish[6] += help
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 6)
	if err != nil {
	log.Fatal(err)
	}
}
