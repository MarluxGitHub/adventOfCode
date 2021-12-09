package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var heightMap[][]int
var basinMap[][]int
var basinSizes[]int


func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()
  linesToHeightMap()
  solve()

  println("")

  sortBasisSizesDesc()
  slice3OfBasinSizes()
  result2 := productOfSlice(basinSizes)

  println(fmt.Sprintf("Result: %v", result2))
}

func sortBasisSizesDesc() {
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
}

func slice3OfBasinSizes() {
	basinSizes = basinSizes[:3]
}

func productOfSlice(slice []int) int {
	product := 1
	for _, v := range slice {
		product *= v
	}
	return product
}

func solve() int {
	sum := 0

	sum += iterateHeightMapMaskSize2()

	return sum
}

func iterateHeightMapMaskSize2() int {
	sum := 0

	println("")
	println("")

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			hit := true
			out:
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					if i+x < 0 || i+x >= len(heightMap) {
						continue
					}
					if j+y < 0 || j+y >= len(heightMap[0]) {
						continue
					}

					if(heightMap[i][j] > heightMap[i+x][j+y]) {
						hit = false
						break out;
					}
				}
			}

			if hit {
				sum += heightMap[i][j] + 1

				growAt(i, j)
				basinSizes = append(basinSizes, getElementCountBiggerThenMinusOneOnBasinMap())
				resetbasinMap()
			}
		}
	}

	return sum
}

func getElementCountBiggerThenMinusOneOnBasinMap() int {
	count := 0
	for i := 0; i < len(basinMap); i++ {
		for j := 0; j < len(basinMap[i]); j++ {
			if basinMap[i][j] == -1 {
				continue
			}

			if basinMap[i][j] >= 0 {
				count++
			}
		}
	}

	return count
}

func resetbasinMap() {
	for i := 0; i < len(basinMap); i++ {
		for j := 0; j < len(basinMap[i]); j++ {
			basinMap[i][j] = -1
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func growAt(i, j int) {

	if i < 0 || i >= len(basinMap) {
		return
	}
	if j < 0 || j >= len(basinMap[i]) {
		return
	}

	if basinMap[i][j] != -1 {
		return
	}

	basinMap[i][j] = heightMap[i][j]

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {

			if(abs(x) + abs(y) == 2) {
				continue
			}

			if x == 0 && y == 0 {
				continue
			}
			if i+x < 0 || i+x >= len(heightMap) {
				continue
			}
			if j+y < 0 || j+y >= len(heightMap[0]) {
				continue
			}

			if(basinMap[i+x][j+y] != -1) {
				continue
			}

			if(heightMap[i+x][j+y] == 9) {
				continue
			}

			if( heightMap[i][j] < heightMap[i+x][j+y]) {
				growAt(i+x, j+y)
			}
		}
	}
}


func linesToHeightMap() {
	heightMap = make([][]int, len(lines))
	basinMap = make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		basinMap[i] = make([]int, len(line))
		for j, c := range line {
			heightMap[i][j] = int(c - '0')
			basinMap[i][j] = -1
		}
	}
}

func printHeightMap(hmap [][]int) {
	for _, row := range hmap {
		println("")
		for _, col := range row {
			printf(fmt.Sprintf("%v ", col))
		}
	}
	println("")
	println("")
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err = i.Strings(2021, 9)
	if err != nil {
	log.Fatal(err)
	}
}
