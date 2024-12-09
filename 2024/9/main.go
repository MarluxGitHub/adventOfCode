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
var year = 2024
var day = 9

var fileIndex []int
var filesystem string
var typedFS []FileSystem

type FileSystem struct {
	size   int
	isFile bool
	file   int
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	generateFileSystem()

	for i := 0; i < len(fileIndex); i++ {
		if fileIndex[i] == -1 {
			for j := len(fileIndex) - 1; j > i; j-- {
				if fileIndex[j] != -1 {
					fileIndex[i], fileIndex[j] = fileIndex[j], fileIndex[i]
					break
				}
			}
		}
	}

	for i := 0; i < len(fileIndex); i++ {
		if fileIndex[i] != -1 {
			result += i * fileIndex[i]
		}
	}
}

// Solve part 2
func Solve2() {
	// TODO Implement my JS solution
}

func generateFileSystem() {
	filesystem = ""
	fileIndex = make([]int, 0)

	var sb strings.Builder
	i := 0
	isFile := true

	for _, c := range lines[0] {
		// convert c to int
		cInt := int(c - '0')

		for range cInt {
			if isFile {
				sb.WriteString(strconv.Itoa(i))
				fileIndex = append(fileIndex, i)
			} else {
				fileIndex = append(fileIndex, -1)
				sb.WriteString(".")
			}
		}

		if isFile {
			i++
		}
		isFile = !isFile
	}

	filesystem = sb.String()

	current := fileIndex[0]
	l := 0
	for i := 0; i < len(fileIndex); i++ {
		if current != fileIndex[i] {
			typedFS = append(typedFS, FileSystem{l, current != -1, current})
			current = fileIndex[i]
			l = 1
		} else {
			l++
		}
	}

	typedFS = append(typedFS, FileSystem{l, current != -1, current})
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
