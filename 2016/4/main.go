package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"

	"MarluxGitHub/adventOfCode/pkg/cryptography"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2016
var day = 4

type Room struct {
	checksum string
	sectorId int
	name     string
}

type KeyValue struct {
	Key   string
	Value int
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
	for _, line := range lines {
		room := parseRoom(line)
		if room.isValid() {
			result += room.sectorId
		}
	}

}

// Solve part 2
func Solve2() {
	for _, line := range lines {
		room := parseRoom(line)
		if room.isValid() {
			decrypted := cryptography.Cipher(room.name, room.sectorId)

			println(strconv.Itoa(room.sectorId) + ":" + decrypted)
		}
	}
}

func parseRoom(line string) Room {
	var room Room
	var checksum string
	var sectorId int
	var name string

	splitted := strings.Split(line, "-")
	name = strings.Join(splitted[:len(splitted)-1], "-")

	fmt.Sscanf(splitted[len(splitted)-1], "%d[%s", &sectorId, &checksum)
	checksum = checksum[:len(checksum)-1]
	room = Room{checksum, sectorId, name}

	return room
}

func (r Room) isValid() bool {
	charCount := make(map[rune]int)
	for _, c := range r.name {
		if c != '-' {
			charCount[c]++
		}
	}

	sorted := make([]KeyValue, len(charCount))

	i := 0
	for k, v := range charCount {
		sorted[i] = KeyValue{string(k), v}
		i++
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	checksum := []rune{}
	for i := 0; i < 5; i++ {
		checksum = append(checksum, []rune(sorted[i].Key)...)
	}

	// check if all runes in r.checksum are in checksum
	for _, c := range r.checksum {
		if !strings.Contains(string(checksum), string(c)) {
			return false
		}
	}

	return true
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
