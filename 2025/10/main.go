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
	totalPresses := 0

	println("Start Solve2")
	println("")
	writer.Flush()

	for _, machine := range Machines {
		presses := solvePart2Z3(machine)
		if presses >= 0 {
			println(fmt.Sprintf("Machine: %d presses", presses))
			writer.Flush()
			totalPresses += presses
		}
	}
	result = totalPresses
}

func solvePart2Z3(m Machine) int {
	// Create Z3 context
	ctx := z3.NewContext(nil)
	intSort := ctx.IntSort()

	// Create integer variables for each button (how many times it's pressed)
	buttonVars := make([]z3.Int, len(m.Buttons))
	for i := range m.Buttons {
		buttonVars[i] = ctx.IntConst(fmt.Sprintf("button%d", i))
	}

	// Create a map from counter index to buttons that affect it
	countersToButtons := make(map[int][]z3.Int)
	for i, button := range m.Buttons {
		for _, flip := range button {
			countersToButtons[flip] = append(countersToButtons[flip], buttonVars[i])
		}
	}

	// Helper function to check if a solution exists with total presses <= maxTotal
	checkSolution := func(maxTotal int) bool {
		solver := z3.NewSolver(ctx)

		// For each counter, add constraint: sum of button presses affecting it = joltage requirement
		for counterIndex, counterButtons := range countersToButtons {
			if counterIndex >= len(m.Joltage) {
				continue
			}
			targetValue := ctx.FromInt(int64(m.Joltage[counterIndex]), intSort).(z3.Int)

			// Sum all button presses that affect this counter
			var sum z3.Int = ctx.FromInt(0, intSort).(z3.Int)
			for _, buttonVar := range counterButtons {
				sum = sum.Add(buttonVar)
			}

			// Add constraint: sum == targetValue
			solver.Assert(sum.Eq(targetValue))
		}

		// Ensure all button variables are non-negative
		zero := ctx.FromInt(0, intSort).(z3.Int)
		for _, buttonVar := range buttonVars {
			solver.Assert(buttonVar.GE(zero))
		}

		// Sum of all button presses
		var sumOfAllButtonVars z3.Int = ctx.FromInt(0, intSort).(z3.Int)
		for _, buttonVar := range buttonVars {
			sumOfAllButtonVars = sumOfAllButtonVars.Add(buttonVar)
		}

		// Add constraint: total presses <= maxTotal
		maxPresses := ctx.FromInt(int64(maxTotal), intSort).(z3.Int)
		solver.Assert(sumOfAllButtonVars.LE(maxPresses))

		// Check if there's a solution
		sat, err := solver.Check()
		return err == nil && sat
	}

	// Binary search for minimum total presses
	// First, find an upper bound
	upperBound := 1
	for !checkSolution(upperBound) {
		upperBound *= 2
		if upperBound > 100000 {
			println("Problem is UNSATISFIABLE (no solution exists).")
			return -1000000
		}
	}

	// Binary search for minimum
	left, right := 0, upperBound
	result := upperBound
	for left <= right {
		mid := (left + right) / 2
		if checkSolution(mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
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
