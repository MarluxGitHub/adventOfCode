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
var result = 0

type Delim struct{
	opening bool
	matcher rune
	score int
}

func (d *Delim) closes(r rune) bool {
	return (*d).matcher == r
}

type StackElement struct {
	val *Delim
	next *StackElement
}

type Stack struct {
	top *StackElement
}

func (s *Stack) push(d *Delim) {
	(*s).top = &StackElement{d, (*s).top}
}

func (s *Stack) pop() *Delim {
	if (*s).top == nil {
	return nil
	}
	d := (*s).top.val
	(*s).top = (*s).top.next
	return d
}

var delimiters = map[rune]*Delim{
	'{': &Delim{true, '}', 1},
	'(': &Delim{true, ')', 2},
	'<': &Delim{true, '>', 3},
	'[': &Delim{true, ']', 4},
	')': &Delim{false, '(', 3},
	']': &Delim{false, '[', 57},
	'}': &Delim{false, '{', 1197},
	'>': &Delim{false, '<', 25137},
}



func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()

  for i := 0; i < len(lines); i++ {
	 result += solveline(i)
  }

  println("Result: " + strconv.Itoa(result))
}

func solveline(i int) int {
	runes := []rune(lines[i])

	stack := &Stack{}

	for _, r := range runes {
		d := delimiters[r]

		if(d.opening) {
			stack.push(d)
		} else {
			p := stack.pop()
			if p == nil {
				break
			} else {
				if(!p.closes(r)) {
					return d.score
				}
			}
		}
	}
	return 0
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 10)
	if err != nil {
	log.Fatal(err)
	}
}
