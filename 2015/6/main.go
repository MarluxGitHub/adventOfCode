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
var year = 2015
var day = 6

func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

  	solve1()
  	println("1:" + strconv.Itoa(result))

  	solve2()
  	println("2:" + strconv.Itoa(result))
}

func solve1() {
	result = 0

	lights := [1000][1000]bool{}

	for _, line := range lines {
		var x1, y1, x2, y2 int
		var action string

		// Get first part of string and check if it is "turn on", "turn off" or "toggle"
		// Write the string to action
		fmt.Sscanf(line, "%s %d,%d through %d,%d", &action, &x1, &y1, &x2, &y2)

		// Check action
		if action == "turn" {
			onOff := ""
			// Get second part of string
			fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &action, &onOff, &x1, &y1, &x2, &y2)
			action = action + " " + onOff
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch action {
				case "turn on":
					lights[x][y] = true
				case "turn off":
					lights[x][y] = false
				case "toggle":
					lights[x][y] = !lights[x][y]
				}
			}
		}
	}

	// double loop and count lights which are true
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if lights[x][y] {
				result++
			}
		}
	}
	// Solve part 1
}

func solve2() {
  	result = 0

	  result = 0

	  lights := [1000][1000]int{}

	  for _, line := range lines {
		  var x1, y1, x2, y2 int
		  var action string

		  // Get first part of string and check if it is "turn on", "turn off" or "toggle"
		  // Write the string to action
		  fmt.Sscanf(line, "%s %d,%d through %d,%d", &action, &x1, &y1, &x2, &y2)

		  // Check action
		  if action == "turn" {
			  onOff := ""
			  // Get second part of string
			  fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &action, &onOff, &x1, &y1, &x2, &y2)
			  action = action + " " + onOff
		  }

		  for x := x1; x <= x2; x++ {
			  for y := y1; y <= y2; y++ {
				  switch action {
				  case "turn on":
					  lights[x][y]++
				  case "turn off":
					  lights[x][y]--
					  if lights[x][y] < 0 {
						  lights[x][y] = 0
					  }
				  case "toggle":
					  lights[x][y]+=2
				  }
			  }
		  }
	  }

	  // double loop and count lights which are true
	  for x := 0; x < 1000; x++ {
		  for y := 0; y < 1000; y++ {
			  result += lights[x][y]
		  }
	  }

	// Solve part 2
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
