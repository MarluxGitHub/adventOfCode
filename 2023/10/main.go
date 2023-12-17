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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 10

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
	m, start := parse()

	s1 := State{start.W(), E}
	s2 := State{start.E(), W}

	var steps int
	for steps = 1; s1.Current != s2.Current; steps++ {
		s1 = s1.Next(m)
		s2 = s2.Next(m)
	}

	result = steps
}

// Solve part 2
func Solve2() {
	m, start := parse()

	mainPipe := map[P]struct{}{start: {}}
	s := State{start.E(), W}
	for s.Current != start {
		mainPipe[s.Current] = struct{}{}
		s = s.Next(m)
	}

	count := 0
	for r := 0; r < 140; r++ {

		inside := false
		temp := ""

		for c := 0; c < 140; c++ {

			p := P{r, c}

			if _, ok := mainPipe[p]; !ok {
				if inside {
					count++
				}
				continue
			}

			switch m[p] {
			case '|':
				inside = !inside
			case 'L', 'F', '7', 'J':
				temp += string(m[p])
				switch temp {
				case "LJ", "F7": // U-shaped, they cancel each other
					temp = ""
				case "L7", "FJ": // like a single vertical line, switch inside
					inside = !inside
					temp = ""
				}
			}
		}
	}

	result = count
}

const (
	N = iota
	E
	S
	W
)

type P struct{ r, c int }

func (p P) N() P { return P{p.r - 1, p.c} }
func (p P) S() P { return P{p.r + 1, p.c} }
func (p P) W() P { return P{p.r, p.c - 1} }
func (p P) E() P { return P{p.r, p.c + 1} }

func parse() (map[P]byte, P) {

	m := map[P]byte{}
	var start P

	for r := 0; r < len(lines); r++ {
		line := lines[r]
		for c := 0; c < len(line); c++ {
			p := P{r, c}

			switch line[c] {
			case 'S':
				start = p
				m[p] = '-'
			case '.':
				continue
			default:
				m[p] = line[c]
			}
		}
	}
	return m, start
}

type State struct {
	Current P
	From    int
}

func (s State) Next(m map[P]byte) State {
	switch m[s.Current] {
	case '|':
		switch s.From {
		case N:
			return State{s.Current.S(), N}
		case S:
			return State{s.Current.N(), S}
		}
	case '-':
		switch s.From {
		case W:
			return State{s.Current.E(), W}
		case E:
			return State{s.Current.W(), E}
		}
	case 'L':
		switch s.From {
		case N:
			return State{s.Current.E(), W}
		case E:
			return State{s.Current.N(), S}
		}
	case 'J':
		switch s.From {
		case N:
			return State{s.Current.W(), E}
		case W:
			return State{s.Current.N(), S}
		}
	case '7':
		switch s.From {
		case S:
			return State{s.Current.W(), E}
		case W:
			return State{s.Current.S(), N}
		}
	case 'F':
		switch s.From {
		case S:
			return State{s.Current.E(), W}
		case E:
			return State{s.Current.S(), N}
		}
	}
	log.Fatal("should never happen")
	return State{}
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
