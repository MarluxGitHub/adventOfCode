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
var result int
var year = 2022
var day = 2

var rpsMapper = map[rune]rune{
	'A': 'R',
	'B': 'P',
	'C': 'S',
	'X': 'R',
	'Y': 'P',
	'Z': 'S',
  }

  var rpsScorer = map[rune]Rps {
	'R': {1, 'S', 'P'},
	'P': {2, 'R', 'S'},
	'S': {3, 'P', 'R'},
}

type Rps struct {
	score int
	beats rune
	loose rune
}


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

	for _, line := range lines {
		enemy :=  rpsMapper[rune (strings.Split(line, " ")[0][0])]
		me :=  rpsMapper[rune (strings.Split(line, " ")[1][0])]

		if rpsScorer[me].beats == enemy {
			result += 6
		}

		if me == enemy {
			result += 3
		}

		result += rpsScorer[me].score
	}
}

func solve2() {
  	result = 0

	for _, line := range lines {
		enemy :=  rpsMapper[rune (strings.Split(line, " ")[0][0])]
		res :=  rune (strings.Split(line, " ")[1][0])
		me := 'A'

		switch res {
			case 'X':
				me = rpsScorer[enemy].beats
			case 'Y':
				me = enemy
			case 'Z':
				me = rpsScorer[enemy].loose
		}

		if rpsScorer[me].beats == enemy {
			result += 6
		}

		if me == enemy {
			result += 3
		}

		result += rpsScorer[me].score
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
