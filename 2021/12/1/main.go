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
var result int = 0

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()
  result := solve()

  println("Result: " + strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 12)
	if err != nil {
	log.Fatal(err)
	}
}

func solve() int {
	m := make(map[string]map[string]bool)
	for _, line := range lines {
		paths := strings.Split(line, "-")
		from, to := paths[0], paths[1]
		if m[from] == nil {
			m[from] = make(map[string]bool)
		}
		if m[to] == nil {
			m[to] = make(map[string]bool)
		}
		m[from][to] = true
		m[to][from] = true
	}

	var count int
	stack := [][]string{{"start"}}
	for len(stack) > 0 {
		var last []string
		stack, last = stack[:len(stack)-1], stack[len(stack)-1]
		tail := last[len(last)-1]
		for k := range m[tail] {
			visited := make(map[string]bool)
			for _, door := range last {
				visited[door] = true
			}
			if strings.ToLower(k) == k && visited[k] {
				continue
			}

			t := make([]string, len(last))
			copy(t, last)
			t = append(t, k)

			if k == "end" {
				count++
			}
			stack = append(stack, t)
		}
	}
	return count
}



