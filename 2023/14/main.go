package main

import (
	"bufio"
	"bytes"
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
var year = 2023
var day = 14

var platform [][]byte

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parsePlatform()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

func parsePlatform() {
	platform = make([][]byte, len(lines))
	for i, line := range lines {
		platform[i] = []byte(line)
	}
}

func Solve1() {
	lenI, lenJ := len(platform), len(platform[0])
	var sum uint64

	for j := 0; j < lenJ; j++ {
		blockI := -1
		ongoing := 0
		for i := 0; i < lenI; i++ {
			switch platform[i][j] {
			case '.':
				continue
			case '#':
				sum += uint64(ongoing*(lenI-blockI-1) - ongoing*(ongoing-1)/2)
				ongoing = 0
				blockI = i
			case 'O':
				ongoing++
			}
		}
		if ongoing > 0 {
			sum += uint64(ongoing*(lenI-blockI-1) - ongoing*(ongoing-1)/2)
		}
	}

	result = int(sum)
}

func Solve2() {
	lenI, lenJ := len(platform), len(platform[0])
	cycle := func() {
		for j := 0; j < lenJ; j++ {
			blockI := -1
			ongoing := 0
			for i := 0; i < lenI; i++ {
				switch platform[i][j] {
				case '.':
					continue
				case '#':
					for n := 0; n < ongoing; n++ {
						platform[blockI+n+1][j] = 'O'
					}
					ongoing = 0
					blockI = i
				case 'O':
					platform[i][j] = '.'
					ongoing++
				}
			}
			for n := 0; n < ongoing; n++ {
				platform[blockI+n+1][j] = 'O'
			}
		}

		for i := 0; i < lenI; i++ {
			blockJ := -1
			ongoing := 0
			for j := 0; j < lenJ; j++ {
				switch platform[i][j] {
				case '.':
					continue
				case '#':
					for n := 0; n < ongoing; n++ {
						platform[i][blockJ+n+1] = 'O'
					}
					ongoing = 0
					blockJ = j
				case 'O':
					platform[i][j] = '.'
					ongoing++
				}
			}
			for n := 0; n < ongoing; n++ {
				platform[i][blockJ+n+1] = 'O'
			}
		}

		for j := 0; j < lenJ; j++ {
			blockI := lenI
			ongoing := 0
			for i := lenI - 1; i >= 0; i-- {
				switch platform[i][j] {
				case '.':
					continue
				case '#':
					for n := 0; n < ongoing; n++ {
						platform[blockI-n-1][j] = 'O'
					}
					ongoing = 0
					blockI = i
				case 'O':
					platform[i][j] = '.'
					ongoing++
				}
			}
			for n := 0; n < ongoing; n++ {
				platform[blockI-n-1][j] = 'O'
			}
		}

		for i := 0; i < lenI; i++ {
			blockJ := lenJ
			ongoing := 0
			for j := lenJ - 1; j >= 0; j-- {
				switch platform[i][j] {
				case '.':
					continue
				case '#':
					for n := 0; n < ongoing; n++ {
						platform[i][blockJ-n-1] = 'O'
					}
					ongoing = 0
					blockJ = j
				case 'O':
					platform[i][j] = '.'
					ongoing++
				}
			}
			for n := 0; n < ongoing; n++ {
				platform[i][blockJ-n-1] = 'O'
			}
		}
	}

	const billion = 1_000_000_000
	breakpoint := billion
	seen := map[string]int{}

	for k := 0; k < billion; k++ {
		cycle()

		if breakpoint == billion {
			h := string(bytes.Join(platform, []byte{}))
			if i, ok := seen[h]; ok {
				breakpoint = k + (billion-i)%(k-i) - 1
			} else {
				seen[h] = k
			}
		}

		if k == breakpoint {
			break
		}

	}

	var sum uint64
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if platform[i][j] == 'O' {
				sum += uint64(lenI - i)
			}
		}
	}

	result = int(sum)
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
