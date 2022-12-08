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
var result int
var year = 2022
var day = 8

type tree struct {
	left *tree
	right *tree
	top *tree
	bottom *tree

	height int
	isVisible bool
	isEdge bool
}

var trees [][]*tree

func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()
	readTreeMap()

	computeNeighborsInTrees()
	setVisibleIfEdgeInTrees()

  	solve1()
  	println("1:" + strconv.Itoa(result))

  	solve2()
  	println("2:" + strconv.Itoa(result))
}

func solve1() {
	result = 0
	// Solve part 1
	computeVisibleInTrees()

	result = getNumberOfVisibleInTrees()
}

func solve2() {
  	result = 0

	computeScores()
  	// Solve part 2
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

func readTreeMap() {
	trees = make([][]*tree, len(lines))

	for i, line := range lines {
		for _, c := range line {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}

			trees[i] = append(trees[i], &tree{height: height, isVisible: false})
		}
	}
}

func computeNeighborsInTrees() {
	for i, line := range trees {
		for j, tree := range line {
			if i > 0 {
				tree.top = trees[i-1][j]
			}
			if i < len(trees)-1 {
				tree.bottom = trees[i+1][j]
			}
			if j > 0 {
				tree.left = trees[i][j-1]
			}
			if j < len(line)-1 {
				tree.right = trees[i][j+1]
			}
		}
	}
}

func setVisibleIfEdgeInTrees() {
	for i, line := range trees {
		for j, tree := range line {
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(line)-1 {
				tree.isVisible = true
				tree.isEdge = true
			}
		}
	}
}

func getNumberOfVisibleInTrees() int {
	numberOfVisible := 0
	for _, line := range trees {
		for _, tree := range line {
			if tree.isVisible {
				numberOfVisible++
			}
		}
	}

	return numberOfVisible
}

func computeVisibleInTrees() {
	for _, line := range trees {
		for _, tree := range line {
			tree.isVisible = isTreeVisible(tree)
		}
	}
}

func isTreeVisible(tree *tree) bool {
	if tree.isVisible {
		return true
	}

	// Check top path
	for t := tree.top; t != nil; t = t.top {
		if t.height >= tree.height {
			break
		}

		if t.isVisible && t.isEdge{
			return true
		}
	}

	// Check bottom path
	for t := tree.bottom; t != nil; t = t.bottom {
		if t.height >= tree.height {
			break
		}

		if t.isVisible && t.isEdge {
			return true
		}
	}

	// Check left path
	for t := tree.left; t != nil; t = t.left {
		if t.height >= tree.height {
			break
		}

		if t.isVisible && t.isEdge {
			return true
		}
	}

	// Check right path
	for t := tree.right; t != nil; t = t.right {
		if t.height >= tree.height {
			break
		}

		if t.isVisible && t.isEdge {
			return true
		}
	}

	return false
}

func computeScores() {
	for _, line := range trees {
		for _, tree := range line {
			result = max(result, computeScore(tree))
		}
	}
}

func computeScore(tree *tree) int {
	topScore := 0
	bottomScore := 0
	leftScore := 0
	rightScore := 0

	// Check top path
	for t := tree.top; t != nil; t = t.top {
		topScore++

		if t.height >= tree.height {
			break
		}
	}

	// Check bottom path
	for t := tree.bottom; t != nil; t = t.bottom {
		bottomScore++

		if t.height >= tree.height {
			break
		}
	}

	// Check left path
	for t := tree.left; t != nil; t = t.left {
		leftScore++

		if t.height >= tree.height {
			break
		}
	}

	// Check right path
	for t := tree.right; t != nil; t = t.right {
		rightScore++

		if t.height >= tree.height {
			break
		}
	}

	return topScore * bottomScore * leftScore * rightScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTreeMap() {
	for _, line := range trees {
		for _, tree := range line {
			if tree.isVisible {
				printf("O")
			} else {
				printf("X")
			}
		}
		println("")
	}
	println("")
}