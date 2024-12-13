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
var year = 2024
var day = 13

type machine struct {
	GetA, GetB             func() int
	ax, ay, bx, by, tx, ty int
}

var machines []machine

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
	genMachines(false)

	for _, m := range machines {
		result += m.GetMinimalTokens()
	}
}

// Solve part 2
func Solve2() {
	genMachines(true)

	for _, m := range machines {
		result += m.GetMinimalTokens()
	}
}

func genMachines(is2 bool) {
	// clear machines
	machines = []machine{}

	for i := 0; i < len(lines); i += 4 {
		// Button A: X+94, Y+34
		var ax, ay, bx, by, tx, ty int

		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bx, &by)

		// Prize: X=8400, Y=5400
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &tx, &ty)

		if is2 {
			tx += 10000000000000
			ty += 10000000000000
		}

		machines = append(machines, machine{
			GetB: func() int {
				return (tx*ay - ty*ax) / (ay*bx - by*ax)
			},

			GetA: func() int {
				return (tx*by - ty*bx) / (by*ax - bx*ay)
			},
			ax: ax, ay: ay, bx: bx, by: by, tx: tx, ty: ty,
		})
	}
}

func (m *machine) GetMinimalTokens() int {
	a := m.GetA()
	b := m.GetB()

	if m.ax*a+m.bx*b == m.tx && m.ay*a+m.by*b == m.ty {
		return 3*a + b
	}

	return 0
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
