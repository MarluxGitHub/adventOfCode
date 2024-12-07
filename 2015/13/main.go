package main

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"MarluxGitHub/adventOfCode/pkg/math"
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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2015
var day = 13

var graph *datastructures.Graph
var nodeMap map[int]*datastructures.Node

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	generateGraph()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	solve()
}

// Solve part 2
func Solve2() {

	graph.AddNode("Me", nil)
	nodeMe := graph.GetNode("Me")
	nodeMap[len(nodeMap)] = nodeMe

	for _, node := range graph.GetNodes() {
		if node.Name != "Me" {
			node.AddEdge(graph.GetNode("Me"), 0)
			nodeMe.AddEdge(node, 0)
		}
	}

	solve()
}

func solve() {
	permutations := math.Permutations(len(graph.GetNodes()))

	for _, permutation := range permutations {
		sum := 0
		for i := 0; i < len(permutation)-1; i++ {
			sum += nodeMap[permutation[i]].GetEdge(nodeMap[permutation[i+1]]).Cost
			sum += nodeMap[permutation[i+1]].GetEdge(nodeMap[permutation[i]]).Cost
		}

		sum += nodeMap[permutation[0]].GetEdge(nodeMap[permutation[len(permutation)-1]]).Cost
		sum += nodeMap[permutation[len(permutation)-1]].GetEdge(nodeMap[permutation[0]]).Cost

		if sum > result {
			result = sum
		}
	}
}

func generateGraph() {
	nodeMap = make(map[int]*datastructures.Node)
	graph = datastructures.NewGraph()

	i := 0
	for _, line := range lines {
		string := strings.Split(line, " ")

		node1 := graph.GetNode(string[0])
		string[10] = string[10][:len(string[10])-1]
		node2 := graph.GetNode(string[10])
		// remove the dot at the end

		weight, _ := strconv.Atoi(string[3])
		if string[2] == "lose" {
			weight = -weight
		}

		if node1 == nil {
			graph.AddNode(string[0], nil)
			node1 = graph.GetNode(string[0])
			nodeMap[i] = node1
			i++
		}

		if node2 == nil {
			graph.AddNode(string[10], nil)
			node2 = graph.GetNode(string[10])
			nodeMap[i] = node2
			i++
		}

		node1.AddEdge(node2, weight)
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
