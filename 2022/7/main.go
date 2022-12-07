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
func printf(f string) { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2022
var day = 7

type directory struct {
	name string
	size int
	subDirs []*directory
	files []*file
	parent *directory
}

type file struct {
	name string
	size int
}

var root directory = directory{name: "/", size: 0, subDirs: []*directory{}, files: []*file{}, parent: nil}

func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()
	readFileSystemFromInput()
	computeDirectorySize(&root)

	result = 0
  	solve1(&root)
  	println("1:" + strconv.Itoa(result))

  	solve2()
  	println("2:" + strconv.Itoa(result))
}

func solve1(dir *directory) {
	for _, dir := range dir.subDirs {
		if(dir.size <= 100000) {
			result += dir.size
		}
		solve1(dir)
	}
}

func solve2() {
	total := 70000000
	result = total
	update := 30000000

	unused := total - root.size
	diff := update - unused

	solve2dfs(&root, diff)
}

func solve2dfs(dir *directory, diff int) {
	for _, dir := range dir.subDirs {
		if(dir.size >= diff) {
			result = min(result, dir.size)
		}
		solve2dfs(dir, diff)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readFileSystemFromInput() {
	current := &root

	lenLines := len(lines)

	for i := 0; i < lenLines; i++ {
		args := strings.Split(lines[i], " ")

		if args[0] == "$" {
			if args[1] == "cd" {
				current = inputCd(current, args[2])
			}

			if args[1] == "ls" {
				i++
				out:
				for j := i; j < lenLines; j++ {
					args = strings.Split(lines[i], " ")

					if args[0] == "$" {
						i--
						break
					}

					switch args[0] {
						case "$":
							i--
							break out
						case "dir":
							subDir := directory{name: args[1], size: 0, subDirs: []*directory{}, files: []*file{}, parent: current}
							current.subDirs = append(current.subDirs, &subDir)
						default:
							size, err := strconv.Atoi(args[0])
							if err != nil {
								log.Fatal(err)
							}
							f := file{name: args[1], size: size}
							current.files = append(current.files, &f)
					}

					i++;
				}
			}
		}
	}
}

func inputCd(current *directory, path string) *directory{
	if path == "/" {
		return &root
	}

	if path == ".." {
		return current.parent
	}

	for _, dir := range current.subDirs {
		if dir.name == path {
			return dir
		}
	}

	return current
}

func computeDirectorySize(dir *directory) int {
	size := 0

	for _, file := range dir.files {
		size += file.size
	}

	for _, subdir := range dir.subDirs {
		size += computeDirectorySize(subdir)
	}

	dir.size = size

	return size
}

func printFile(file file) {
	println(file.name + " " + strconv.Itoa(file.size))
}

func printDir(dir directory) {
	println(dir.name + " " + "subdiroctoryCount " + strconv.Itoa(len(dir.subDirs)) + " " + "fileCount " + strconv.Itoa(len(dir.files)))
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
