package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"MarluxGitHub/adventOfCode/pkg/math"
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
var year = 2015
var day = 9

var graph = datastructures.NewGraph()
var perms [][]int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genGraph()
	perms = math.Permutations(len(graph.GetNodes()))

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	nodes := graph.GetNodes()

	size := len(nodes)

	// get all list of List of all combinations of numbers from 0 to size

	min := int32(999999)

outer:
	for _, perm := range perms {
		sum := int32(0)
		for i := 0; i < size-1; i++ {
			from := nodes[perm[i]]
			to := nodes[perm[i+1]]

			edge := from.GetEdge(to)

			if edge == nil {
				continue outer
			}

			sum += int32(edge.Cost)
		}

		if sum < min {
			min = sum
		}
	}

	result = int(min)

}

// Solve part 2
func Solve2() {
	nodes := graph.GetNodes()

	size := len(nodes)

	// get all list of List of all combinations of numbers from 0 to size

	max := int32(0)

outer:
	for _, perm := range perms {
		sum := int32(0)
		for i := 0; i < size-1; i++ {
			from := nodes[perm[i]]
			to := nodes[perm[i+1]]

			edge := from.GetEdge(to)

			if edge == nil {
				continue outer
			}

			sum += int32(edge.Cost)
		}

		if sum > max {
			max = sum
		}
	}

	result = int(max)
}

func genGraph() {
	for _, i := range lines {
		var from, to string
		var cost int
		fmt.Sscanf(i, "%s to %s = %d", &from, &to, &cost)

		fromNode := graph.GetNode(from)
		if fromNode == nil {
			fromNode = graph.AddNode(from)
		}

		toNode := graph.GetNode(to)
		if toNode == nil {
			toNode = graph.AddNode(to)

		}

		fromNode.AddEdge(toNode, cost)
		toNode.AddEdge(fromNode, cost)
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
