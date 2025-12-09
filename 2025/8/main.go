package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 8

type JunctionBox struct {
	Position datastructures.Vector
}

var JunctionBoxes []JunctionBox

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genJunctionBoxes()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func genJunctionBoxes() {
	JunctionBoxes = make([]JunctionBox, len(lines))
	for i, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		JunctionBoxes[i] = JunctionBox{Position: datastructures.Vector{X: x, Y: y, Z: z}}
	}
}

// Solve part 1
func Solve1() {
	// parse junction box coordinates already done in genJunctionBoxes
	n := len(JunctionBoxes)
	if n == 0 {
		result = 0
		return
	}

	type pair struct {
		i, j int
		dist int64
	}

	pairs := make([]pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		xi := JunctionBoxes[i].Position.X
		yi := JunctionBoxes[i].Position.Y
		zi := JunctionBoxes[i].Position.Z
		for j := i + 1; j < n; j++ {
			xj := JunctionBoxes[j].Position.X
			yj := JunctionBoxes[j].Position.Y
			zj := JunctionBoxes[j].Position.Z
			dx := int64(xi - xj)
			dy := int64(yi - yj)
			dz := int64(zi - zj)
			d2 := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, pair{i: i, j: j, dist: d2})
		}
	}

	sort.Slice(pairs, func(a, b int) bool { return pairs[a].dist < pairs[b].dist })

	// disjoint set union (union by size)
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) bool {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return false
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
		return true
	}

	limit := 1000
	if limit > len(pairs) {
		limit = len(pairs)
	}
	for k := 0; k < limit; k++ {
		p := pairs[k]
		union(p.i, p.j)
	}

	// compute component sizes
	compSize := make(map[int]int)
	for i := 0; i < n; i++ {
		r := find(i)
		compSize[r]++
	}

	sizes := make([]int, 0, len(compSize))
	for _, s := range compSize {
		sizes = append(sizes, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	prod := 1
	for i := 0; i < 3; i++ {
		if i < len(sizes) {
			prod *= sizes[i]
		} else {
			prod *= 1
		}
	}
	result = prod
}

// Solve part 2
func Solve2() {
	// connect pairs until all boxes are in one component
	n := len(JunctionBoxes)
	if n <= 1 {
		result = 0
		return
	}

	type pair struct {
		i, j int
		dist int64
	}
	pairs := make([]pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		xi := JunctionBoxes[i].Position.X
		yi := JunctionBoxes[i].Position.Y
		zi := JunctionBoxes[i].Position.Z
		for j := i + 1; j < n; j++ {
			xj := JunctionBoxes[j].Position.X
			yj := JunctionBoxes[j].Position.Y
			zj := JunctionBoxes[j].Position.Z
			dx := int64(xi - xj)
			dy := int64(yi - yj)
			dz := int64(zi - zj)
			d2 := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, pair{i: i, j: j, dist: d2})
		}
	}

	sort.Slice(pairs, func(a, b int) bool { return pairs[a].dist < pairs[b].dist })

	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}
	union := func(a, b int) bool {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return false
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
		return true
	}

	components := n
	lastProd := 0
	for _, p := range pairs {
		if union(p.i, p.j) {
			components--
			if components == 1 {
				x1 := JunctionBoxes[p.i].Position.X
				x2 := JunctionBoxes[p.j].Position.X
				lastProd = x1 * x2
				break
			}
		}
	}

	result = lastProd
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
