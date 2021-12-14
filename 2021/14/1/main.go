package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int = 0

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()

  templateStr, polymerMap := inputTransform()


  result = solve(templateStr, polymerMap, 10)

  println(strconv.Itoa(result))
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 14)
	if err != nil {
	log.Fatal(err)
	}
}

func inputTransform() (string, map[string]string) {
	templateStr := ""
	polymerMap := make(map[string]string)

	for _, line := range lines {

		if len(line) <= 0 {
			continue
		}

		if strings.Contains(line, "->") {
			curStr := strings.Split(line, " -> ")
			key := curStr[0]
			val := curStr[1]

			polymerMap[key] = val
		} else {
			templateStr = line
		}
	}

	return templateStr, polymerMap
}

func solve(templateStr string, polymerMap map[string]string, x int) int {
	charMap := make(map[string]int)

	for i := 0; i < len(templateStr)-1; i++ {
		charMap[templateStr[i:i+2]] += 1
	}
	charMap[templateStr[len(templateStr)-1:]] += 1

	for i := 0; i < x; i++ {
		newCharMap := make(map[string]int)

		for k, v := range charMap {
			if polymerMap[k] != "" {
				newCharMap[k[0:1]+polymerMap[k]] += v
				newCharMap[polymerMap[k]+k[1:2]] += v
			} else {
				newCharMap[k] += v
			}
		}
		charMap = newCharMap
	}

	freq := make(map[string]int)
	for k, v := range charMap {
		freq[k[0:1]] += v
	}

	minVal := math.MaxInt64
	maxVal := 0
	minChar := ""
	maxChar := ""
	for key, val := range freq {
		if val < minVal {
			minVal = val
			minChar = key
		}

		if val > maxVal {
			maxVal = val
			maxChar = key
		}
	}

	fmt.Printf("MaxChar: %v MaxVal: %d ----- MinChar: %v MinVal: %d\n", maxChar, maxVal, minChar, minVal)

	return maxVal - minVal
}
