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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2025
var day = 11

type Server struct {
	Outputs []string
}

var servers map[string]Server

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	genServers()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func genServers() {
	servers = make(map[string]Server)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		device := strings.TrimSpace(parts[0])
		outputsStr := strings.TrimSpace(parts[1])
		if outputsStr == "" {
			servers[device] = Server{Outputs: []string{}}
		} else {
			outputs := strings.Fields(outputsStr)
			servers[device] = Server{Outputs: outputs}
		}
	}
}

// Solve part 1
func Solve1() {
	genServers()
	result = countPaths("you", "out")
}

// countPaths counts all paths from start to end using DFS
func countPaths(start, end string) int {
	if start == end {
		return 1
	}

	server, exists := servers[start]
	if !exists {
		return 0
	}

	count := 0
	for _, output := range server.Outputs {
		count += countPaths(output, end)
	}

	return count
}

// Solve part 2
func Solve2() {
	result = countPathsWithRequiredNodes("svr", "out", []string{"dac", "fft"})
}

// countPathsWithRequiredNodes counts all paths from start to end that visit all required nodes
func countPathsWithRequiredNodes(start, end string, required []string) int {
	// Create a map for required nodes for quick lookup
	requiredMap := make(map[string]bool)
	for _, req := range required {
		requiredMap[req] = true
	}

	// Memoization: key is "node|visited1,visited2,..." where visited are the required nodes visited so far
	memo := make(map[string]int)

	// Use a set to track visited nodes in current path (to avoid cycles)
	pathVisited := make(map[string]bool)

	// Track which required nodes have been visited (as a set for memo key)
	visitedRequired := make(map[string]bool)

	return countPathsWithRequiredDFS(start, end, requiredMap, pathVisited, visitedRequired, memo)
}

// buildMemoKey creates a memoization key from current node and visited required nodes
func buildMemoKey(current string, requiredMap map[string]bool, visitedRequired map[string]bool) string {
	key := current + "|"
	for req := range requiredMap {
		if visitedRequired[req] {
			key += req + ","
		}
	}
	return key
}

// allRequiredVisited checks if all required nodes have been visited
func allRequiredVisited(requiredMap map[string]bool, visitedRequired map[string]bool) bool {
	for req := range requiredMap {
		if !visitedRequired[req] {
			return false
		}
	}
	return true
}

// countPathsWithRequiredDFS counts paths using DFS with memoization
func countPathsWithRequiredDFS(current, end string, requiredMap map[string]bool, pathVisited map[string]bool, visitedRequired map[string]bool, memo map[string]int) int {
	if current == end {
		if allRequiredVisited(requiredMap, visitedRequired) {
			return 1
		}
		return 0
	}

	inPath := pathVisited[current]
	memoKey := ""
	if !inPath {
		memoKey = buildMemoKey(current, requiredMap, visitedRequired)
		if count, exists := memo[memoKey]; exists {
			return count
		}
	}

	pathVisited[current] = true
	wasRequiredVisited := visitedRequired[current]
	if requiredMap[current] {
		visitedRequired[current] = true
	}

	defer func() {
		delete(pathVisited, current)
		if requiredMap[current] {
			visitedRequired[current] = wasRequiredVisited
		}
	}()

	server, exists := servers[current]
	if !exists {
		return 0
	}

	count := 0
	for _, output := range server.Outputs {
		if pathVisited[output] {
			continue
		}
		count += countPathsWithRequiredDFS(output, end, requiredMap, pathVisited, visitedRequired, memo)
	}

	if memoKey != "" {
		memo[memoKey] = count
	}

	return count
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
