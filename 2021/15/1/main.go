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
var result int = 0

var problem [][] int
var dp [][] int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()
	readInput()
	stringLinesTo2dIntArray()
	result := solve()

	println(strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 15)
	if err != nil {
	log.Fatal(err)
	}
}

func stringLinesTo2dIntArray () {
	problem = make([][]int, len(lines))
	dp = make([][]int, len(lines))

	for i, line := range lines {
		problem[i] = make([]int, len(line))
		dp[i] = make([]int, len(line))

		for j, char := range line {
			problem[i][j] = int(char) - 48
			dp[i][j] = 0
		}
	}
}

func solve() int {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + problem[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + problem[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + problem[i][j]
			}
		}
	}

	return dp[len(lines)-1][len(lines[0])-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
