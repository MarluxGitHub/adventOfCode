package main

import (
	luxMath "MarluxGitHub/adventOfCode/pkg/math"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 5

type GardenConv struct {
	From, To  string
	Converter func(int) int
}

var ConverterMap = map[string]*GardenConv{}
var Seeds = []int{}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseInput()

	result1 := make(chan int)
	result2 := make(chan int)

	go func() {
		result := 0
		Solve1(&result)
		result1 <- result
	}()

	go func() {
		result := 0
		Solve2(&result)
		result2 <- result
	}()

	println("1:" + strconv.Itoa(<-result1))
	println("2:" + strconv.Itoa(<-result2))
}

func parseInput() {
	parseSeeds()

	from := 2
	for i, line := range lines[2:] {
		if strings.Contains(line, "map:") {
			from = i + 2
			continue
		}

		if strings.Trim(line, " ") == "" {
			parseConverter(from, i+2)
			from = 2
		}
	}

	parseConverter(from, len(lines)-1)
}

func parseSeeds() {
	line := lines[0]

	seedsElements := strings.Split(strings.Split(line, ":")[1], " ")

	for _, seedElement := range seedsElements {
		if seedElement == "" {
			continue
		}

		seed, err := strconv.Atoi(seedElement)

		if err != nil {
			log.Fatal(err)
		}

		Seeds = append(Seeds, seed)
	}
}

func parseConverter(from, to int) {
	// x-to-y map:
	// parse x and y from line[from]

	var x, y string

	l := strings.Replace(lines[from], "-", " ", -1)
	fmt.Sscanf(l, "%s to %s map:", &x, &y)

	converterLines := [][]int{}
	for i := from + 1; i < to; i++ {
		converterLines = append(converterLines, []int{})
	}

	// parse converter from line[from+1:to]
	for i := from + 1; i < to; i++ {
		line := lines[i]
		elements := strings.Split(line, " ")

		for _, element := range elements {
			if element == "" {
				continue
			}

			n, err := strconv.Atoi(element)

			if err != nil {
				log.Fatal(err)
			}

			converterLines[i-(from+1)] = append(converterLines[i-(from+1)], n)
		}
	}

	// Sort converterLines for the second element
	sort.Slice(converterLines, func(i, j int) bool {
		return converterLines[i][1] < converterLines[j][1]
	})

	// Create converter function
	converter := func(n int) int {
		for _, line := range converterLines {
			if n >= line[1] && n < line[1]+line[2] {
				return (line[0] - line[1]) + n
			}
		}

		return n
	}

	ConverterMap[x] = &GardenConv{
		From:      x,
		To:        y,
		Converter: converter,
	}
}

// Solve part 1
func Solve1(result *int) {
	min := math.MaxInt

	for _, seed := range Seeds {
		current := "seed"
		for ConverterMap[current] != nil {
			conv := ConverterMap[current]
			seed = conv.Converter(seed)
			current = conv.To
		}

		min = luxMath.Min(min, seed)
	}

	*result = min
}

// Solve part 2
func Solve2(result *int) {
	min := math.MaxInt
	minValues := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(Seeds); i += 2 {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localMin := math.MaxInt64
			for j := start; j < start+end; j++ {
				current := "seed"
				currentValue := j
				for ConverterMap[current] != nil {
					conv := ConverterMap[current]
					currentValue = conv.Converter(currentValue)
					current = conv.To
				}

				localMin = luxMath.Min(localMin, currentValue)
			}
			minValues <- localMin
		}(Seeds[i], Seeds[i+1])
	}

	go func() {
		wg.Wait()
		close(minValues)
	}()

	for localMin := range minValues {
		min = luxMath.Min(min, localMin)
	}

	*result = min
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
