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
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2016
var day = 10

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

type Bot struct {
	id                int
	low, high         int
	LowWork, HighWork func(int)
}

func newBot(id int) *Bot {
	return &Bot{id: id, low: -1, high: -1}
}

func (b *Bot) addChip(chip int) {
	if b.low == -1 {
		b.low = chip
	} else if b.high == -1 {
		b.high = chip
		if b.low > b.high {
			b.low, b.high = b.high, b.low
		}
		b.compare()
	} else {
		log.Fatal("Bot already has two chips")
	}
}

func (b *Bot) compare() {
	if b.low == 17 && b.high == 61 {
		result = b.id
	}

	b.LowWork(b.low)
	b.HighWork(b.high)

	b.low = -1
	b.high = -1
}

var output = make(map[int]*int)
var bots = make(map[int]*Bot)

// Solve part 1
func Solve1() {
	for _, line := range lines {
		parseLineBot(line)
	}

	for _, line := range lines {
		parseLineVal(line)
	}

}

// Solve part 2
func Solve2() {
	result = *output[0] * *output[1] * *output[2]
}

func parseLineBot(line string) {
	instruction := line[:3]

	switch instruction {
	case "bot":
		readBotLine(line)
	default:
	}
}

func parseLineVal(line string) {
	instruction := line[:3]

	switch instruction {
	case "val":
		readValLine(line)
	default:
	}
}

func readBotLine(line string) {
	var botnumber, low, high int
	var lowType, highType string

	fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &botnumber, &lowType, &low, &highType, &high)

	if _, ok := bots[botnumber]; !ok {
		bots[botnumber] = newBot(botnumber)
	}

	if lowType == "bot" {
		bots[botnumber].LowWork = func(i int) { bots[low].addChip(i) }
	} else {
		bots[botnumber].LowWork = func(i int) { output[low] = &i }
	}

	if highType == "bot" {
		bots[botnumber].HighWork = func(i int) { bots[high].addChip(i) }
	} else {
		bots[botnumber].HighWork = func(i int) { output[high] = &i }
	}
}

func readValLine(line string) {
	var chip, botnumber int

	fmt.Sscanf(line, "value %d goes to bot %d", &chip, &botnumber)

	if _, ok := bots[botnumber]; !ok {
		bots[botnumber] = newBot(botnumber)
	}

	bots[botnumber].addChip(chip)
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
