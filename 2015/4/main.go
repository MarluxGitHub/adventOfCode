package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
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
var year = 2015
var day = 4


func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()
	readInput()
	solve1()
	println("1:" + strconv.Itoa(result))
	solve2()
	println("2:" + strconv.Itoa(result))
}

func solve1() {
	result = 0

	for _, line := range lines {
		for i := 0; ; i++ {
			if generateMd5(line + strconv.Itoa(i))[:5] == "00000" {
				result = i
				break
			}
		}
	}
	// Solve part 1
}

func solve2() {

	for _, line := range lines {
		for i := result; ; i++ {
			if generateMd5(line + strconv.Itoa(i))[:6] == "000000" {
				result = i
				break
			}
		}
	}
}

func generateMd5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
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
