package main

import (
	"bufio"
	"crypto/md5"
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
var result string
var result2 string
var year = 2016
var day = 5


func main() {
  	// STDOUT MUST BE FLUSHED MANUALLY!!!
  	defer writer.Flush()

  	readInput()

	result = ""
	result2 = ""
  	Solve()

  	println("1:" + result)
  	println("2:" + result2)
}

// Solve part 1
func Solve() {

	doorId := lines[0]

	times, times2, i := 0, 0,0


	res2 := make([]string, 8)

	for times2 < 8 {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(doorId + strconv.Itoa(i))))
		if hash[:5] == "00000" {
			if times < 8 {
				result += string(hash[5])
			}

			pos := string(hash[5])
			posInt, err := strconv.Atoi(pos)

			if err != nil {
				continue
			}

			if posInt == 0  {
				res2[posInt] = string(hash[6])
				times2++
			}

			times++
		}
		i++
	}

	result2 = fmt.Sprintf("%v", res2)

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
