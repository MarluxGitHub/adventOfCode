package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 8

type Path struct {
	Label       string
	Left, Right *Path
}

var Labyrinth = map[string]*Path{}
var Instructions string
var StartingPoints = []string{}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseInput()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	current := "AAA"

	for {
		for i := 0; i < len(Instructions); i++ {
			switch Instructions[i] {
			case 'L':
				current = Labyrinth[current].Left.Label
			case 'R':
				current = Labyrinth[current].Right.Label
			}

			result++
			if current == "ZZZ" {
				return
			}
		}
	}
}

// Solve part 2
func Solve2() {
	for _, p := range Labyrinth {
		// Check if last symbol of Label is an A
		if p.Label[len(p.Label)-1] == 'A' {
			StartingPoints = append(StartingPoints, p.Label)
		}
	}

	steps := make([]int, 0)

	for _, p := range StartingPoints {
		step := 0

	out:
		for {
			for j := 0; j < len(Instructions); j++ {
				switch Instructions[j] {
				case 'L':
					p = Labyrinth[p].Left.Label
				case 'R':
					p = Labyrinth[p].Right.Label
				}

				step++

				if p[len(p)-1] == 'Z' {
					steps = append(steps, step)
					break out
				}
			}
		}
	}

	factors := make([]int, 0)
	for _, s := range steps {
		for _, f := range PrimeFactors(s) {
			if !slices.Contains(factors, f) {
				factors = append(factors, f)
			}
		}
	}

	result = 1
	for _, f := range factors {
		result *= f
	}
}

func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func parseInput() {
	Instructions = lines[0]

	for i := 2; i < len(lines); i++ {
		// xxx = (yyy, zzz)

		pattern := `(\w+)\s=\s\((\w+),\s(\w+)\)`
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(lines[i])

		if len(match) > 0 {
			label := match[1]
			left := match[2]
			right := match[3]

			if _, ok := Labyrinth[label]; !ok {
				Labyrinth[label] = &Path{Label: label}
			}

			if _, ok := Labyrinth[left]; !ok {
				Labyrinth[left] = &Path{Label: left}
			}

			if _, ok := Labyrinth[right]; !ok {
				Labyrinth[right] = &Path{Label: right}
			}

			Labyrinth[label].Left = Labyrinth[left]
			Labyrinth[label].Right = Labyrinth[right]
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
