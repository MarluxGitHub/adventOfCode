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

	result = solve()

	print2dArray(problem)

	println("")

	print2dArray(dp)

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
	problem = make([][]int, 5*len(lines))
	dp = make([][]int, 5*len(lines))

	xLen := len(lines)
	yLen := len(lines[0])

	for i, line := range lines {

		for x:=0; x < 5; x++ {
			problem[i + (xLen*x)] = make([]int, 5*len(line))
			dp[i + (xLen*x)] = make([]int, 5*len(line))
		}

		for j, char := range line {
			number := int(char) - 48

			for x:= 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					n := number + x + y
					if(n > 9) {
						n = n - 9
					}

					dp[i+(xLen * x)][j+(yLen * y)] = 0
					problem[i+(xLen * x)][j+(yLen * y)] = n
				}
			}
		}
	}
}

func solve() int {
	for i := 0; i < len(problem); i++ {
		for j := 0; j < len(problem[i]); j++ {
			if j == 0 && i == 0 {
				dp[i][j] = 0
				continue
			}

			if(j == 0) {
				dp[i][j] = problem[i][j] + dp[i-1][j]
				continue
			}

			if(i == 0) {
				dp[i][j] = problem[i][j] + dp[i][j-1]
				continue
			}

			dp[i][j] = problem[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[len(problem) - 1][len(problem[0]) - 1]
}

func print2dArray(array [][]int) {
	for _, line := range array {
		for _, char := range line {
			printf(fmt.Sprintf("%d", char))
			printf(" ")
		}
		println("")
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}
