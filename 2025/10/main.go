package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aclements/go-z3/z3"
	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 10

type Machine struct {
	Indicator uint64
	Buttons   [][]int
	Joltage   []int
}

var Machines []Machine

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genMachines()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func parseInts(s string) []int {
	var result []int
	parts := strings.FieldsFunc(s, func(r rune) bool { return r < '0' || r > '9' })
	for _, p := range parts {
		if n, err := strconv.Atoi(p); err == nil {
			result = append(result, n)
		}
	}
	return result
}

func genMachines() {
	Machines = make([]Machine, len(lines))
	for i, line := range lines {
		// Pattern: [indicator] (button1) (button2) ... {joltage}
		// Extract indicator from [...]
		reBrackets := regexp.MustCompile(`\[([^\]]+)\]`)
		mBrackets := reBrackets.FindStringSubmatch(line)
		var indicator uint64 = 0
		if len(mBrackets) > 1 {
			for idx, ch := range mBrackets[1] {
				if ch == '#' {
					indicator |= (1 << uint(idx))
				}
			}
		}

		// Extract buttons from (...)
		reParen := regexp.MustCompile(`\(([^)]+)\)`)
		mParens := reParen.FindAllStringSubmatch(line, -1)
		var buttons [][]int
		for _, m := range mParens {
			buttons = append(buttons, parseInts(m[1]))
		}

		// Extract joltage from {...}
		reBraces := regexp.MustCompile(`\{([^}]+)\}`)
		mBraces := reBraces.FindStringSubmatch(line)
		var joltage []int
		if len(mBraces) > 1 {
			joltage = parseInts(mBraces[1])
		}

		Machines[i] = Machine{
			Indicator: indicator,
			Buttons:   buttons,
			Joltage:   joltage,
		}
	}
}

// Solve part 1
func Solve1() {
	totalPresses := 0

	println("Start Solve1")
	println("")
	writer.Flush()

	for _, machine := range Machines {
		presses := minButtonPressesBFS(machine)
		if presses >= 0 {
			println(fmt.Sprintf("Machine: %d presses", presses))
			writer.Flush()
			totalPresses += presses
		}
	}
	result = totalPresses
}

// minButtonPressesBFS uses BFS to find minimum button presses (fewest presses to reach target)
func minButtonPressesBFS(m Machine) int {
	// Determine number of lights
	numLights := 0
	for i := 0; i < 64; i++ {
		if (m.Indicator>>uint(i))&1 == 1 {
			numLights = i + 1
		}
	}
	for _, btn := range m.Buttons {
		for _, light := range btn {
			if light+1 > numLights {
				numLights = light + 1
			}
		}
	}

	target := m.Indicator

	// BFS: state = current light config, cost = number of button presses
	type state struct {
		lights  uint64
		presses int
	}

	visited := make(map[uint64]bool)
	queue := make([]state, 0)
	queue = append(queue, state{0, 0})
	visited[0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.lights == target {
			return cur.presses
		}

		// Pruning: avoid exploring too deep
		if cur.presses > 20 {
			continue
		}

		// Try pressing each button once
		for _, btn := range m.Buttons {
			nextLights := cur.lights
			for _, light := range btn {
				if light < numLights {
					nextLights ^= (1 << uint(light))
				}
			}

			if !visited[nextLights] {
				visited[nextLights] = true
				queue = append(queue, state{nextLights, cur.presses + 1})
			}
		}
	}

	return -1 // No solution found
}

// Solve part 2
func Solve2() {
	println("Start Solve2")
	writer.Flush()
	totalPresses := 0
	for _, machine := range Machines {
		presses := solveMachine2(machine.Joltage, machine.Buttons)
		if presses >= 0 {
			println(fmt.Sprintf("Machine: %d presses", presses))
			writer.Flush()
			totalPresses += presses
		}
	}
	result = totalPresses
}

// solveMachine2 - use greedy approach: always press buttons with most effect
func solveMachine2(target []int, buttons [][]int) int {
	state := make([]int, len(target))
	copy(state, target)

	presses := 0
	maxIter := 10000

	for presses < maxIter {
		// Check if done
		done := true
		for _, v := range state {
			if v > 0 {
				done = false
				break
			}
		}
		if done {
			return presses
		}

		// Find button that reduces max value
		bestBtn := -1
		bestReduction := 0

		for btnIdx, btn := range buttons {
			// Check if this button would help
			reduction := 0
			for _, idx := range btn {
				if idx < len(state) && state[idx] > 0 {
					reduction++
				}
			}
			if reduction > bestReduction {
				bestReduction = reduction
				bestBtn = btnIdx
			}
		}

		if bestBtn == -1 {
			break
		}

		// Apply best button
		for _, idx := range buttons[bestBtn] {
			if idx < len(state) {
				state[idx]--
			}
		}
		presses++
	}

	return -1
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
