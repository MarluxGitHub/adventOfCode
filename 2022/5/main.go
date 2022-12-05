package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var stack [][]string

var size int
var result string
var year = 2022
var day = 5


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

  	solve1()
  	println("1:" + result)

  	solve2()
  	println("2:" + result)
}

func solve1() {
	result = ""

	readStack()

	for i := size+1; i < len(lines); i++ {
		//print i

		println(lines[i])

		var count, from, to int = 0, 0, 0

		fmt.Sscanf(lines[i], "move %d from %d to %d", &count, &from, &to)
		from--
		to--

		// // Move count elements from stack[from] to stack[to]
		for j := 0; j < count; j++ {
			stack[to] = append(stack[to], stack[from][len(stack[from])-1])
			stack[from] = stack[from][:len(stack[from])-1]
		}

		printStack()
	}

	// Get Last Element of Every Stack
	for _, s := range stack {
		if(len(s) > 0) {
			result += s[len(s)-1][1:2]
		}
	}
}

func solve2() {
  	result = ""

	readStack()

	for i := size+1; i < len(lines); i++ {
		//print i

		println(lines[i])

		var count, from, to int = 0, 0, 0

		fmt.Sscanf(lines[i], "move %d from %d to %d", &count, &from, &to)
		from--
		to--

		movingStack := []string{}
		for j := 0; j < count; j++ {
			movingStack = append(movingStack, stack[from][len(stack[from])-1])
			stack[from] = stack[from][:len(stack[from])-1]
		}

		for j := 0; j < count; j++ {
			stack[to] = append(stack[to], movingStack[len(movingStack)-1])
			movingStack = movingStack[:len(movingStack)-1]
		}

		printStack()
	}

	// Get Last Element of Every Stack
	for _, s := range stack {
		if(len(s) > 0) {
			result += s[len(s)-1][1:2]
		}
	}
}

func readStack() {
	// Find out how many lines are in the stack
	stackSize := 0

	for _, line := range lines {
		if line == "" {
			break
		}

		stackSize++
	}

	size = stackSize
	stackSize--

	// Get Last Character of line[stackSize] as int
	stackCount := lines[stackSize][len(lines[stackSize])-1]

	// Create StackCount stacks
	stack = make([][]string, stackCount)

	for i := stackSize-1; i >= 0; i-- {
		x := 0
		// split line every 3 characters
		for j := 0; j < len(lines[i]); j += 3 {
			elem := lines[i][j:j+3]
			if elem != "   " {
				stack[x] = append(stack[x], lines[i][j:j+3])
			}
			x++
			j++
		}
	}

	printStack()
}

func printStack(){
	for _, s := range stack {
		somethingPrinted := false
		for _, e := range s {
			printf(e)
			somethingPrinted = true
		}
		if somethingPrinted {
			println("")
		}
	}
	println("")
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
