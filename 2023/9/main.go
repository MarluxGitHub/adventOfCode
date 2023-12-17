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
var year = 2023
var day = 9


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()
	ParseSequence()

	result = 0
  	Solve1()
  	println("1:" + strconv.Itoa(result))

	result = 0
  	Solve2()
  	println("2:" + strconv.Itoa(result))
}

type Sequence struct {
	Number map[int][]int
}

var Sequences = make(map[int]Sequence)

func (s Sequence) ComputeSubSequences(i int) {
	// Check if current Sequence has only zeros
	// If so, return
	onlyZeroes := true
	for _, n := range s.Number[i] {
		if n != 0 {
			onlyZeroes = false
			break
		}
	}

	if onlyZeroes {
		return
	}

	s.Number[i+1] = make([]int, len(s.Number[i])-1)

	for j := 0; j < len(s.Number[i])-1; j++ {
		s.Number[i+1][j] = s.Number[i][j+1] - s.Number[i][j]
	}

	s.ComputeSubSequences(i+1)
}

func (s Sequence) ExtropolateSubSequencesEnd() {
	current := len(s.Number)-1

	s.Number[current] = append(s.Number[current], 0)

	current--

	for current >= 0 {
		newNumber := s.Number[current][len(s.Number[current])-1] + s.Number[current+1][len(s.Number[current+1])-1]
		s.Number[current] = append(s.Number[current], newNumber)
		current--
	}
}

func (s Sequence) ExtropolateSubsequencesStart() {
	current := len(s.Number)-1

	s.Number[current] = append([]int{0}, s.Number[current]...)

	current--

	for current >= 0 {
		newNumber := s.Number[current][0] - s.Number[current+1][0]
		s.Number[current] = append([]int{newNumber}, s.Number[current]...)
		current--
	}
}

// Solve part 1
func Solve1() {
	for _, s := range Sequences {
		s.ComputeSubSequences(0)
		s.ExtropolateSubSequencesEnd()
		result += s.Number[0][len(s.Number[0])-1]
	}

}

// Solve part 2
func Solve2() {
	for _, s := range Sequences {
		s.ComputeSubSequences(0)
		s.ExtropolateSubsequencesStart()
		result += s.Number[0][0]
	}
}

func ParseSequence() {
	for i, line := range lines {
		numbers := strings.Split(line, " ")
		numbersInt := make([]int, len(numbers))
		for i, n := range numbers {
			numbersInt[i], _ = strconv.Atoi(n)
		}
		Sequences[i] = Sequence{
			Number: make(map[int][]int),
		}

		Sequences[i].Number[0] = numbersInt
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
