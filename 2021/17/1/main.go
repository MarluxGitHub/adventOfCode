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

var lines []string
var result int = 0

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()
	readInput()

	result = solve(217,240,-126,-69)

	println(strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err = i.Strings(2021, 17)
	if err != nil {
		log.Fatal(err)
	}
}

func solve(x1, x2, y1, y2 int) int {
	return (y1 * (y1+1))/2
}