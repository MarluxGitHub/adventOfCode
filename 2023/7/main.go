package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 7

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

type Hand struct {
	hand string
	bid  int
}

func handType(hand Hand, part2 bool) int {
	m := map[byte]int{}
	for i := range hand.hand {
		m[hand.hand[i]]++
	}

	// sort cards by count
	type card struct {
		c   byte
		cnt int
	}
	cards := []card{}
	for k, v := range m {
		cards = append(cards, card{k, v})
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].cnt > cards[j].cnt
	})

	if part2 {
		cnt := 0
		if len(cards) > 1 {
			for i := range cards {
				if cards[i].c == 'J' {
					cnt = cards[i].cnt
					cards = append(cards[:i], cards[i+1:]...) // remove 'J'
					break
				}
			}
			cards[0].cnt += cnt
		}
	}

	if cards[0].cnt == 5 {
		return 6
	} else if cards[0].cnt == 4 {
		return 5
	} else if cards[0].cnt == 3 && cards[1].cnt == 2 {
		return 4
	} else if cards[0].cnt == 3 && cards[1].cnt == 1 {
		return 3
	} else if cards[0].cnt == 2 && cards[1].cnt == 2 {
		return 2
	} else if cards[0].cnt == 2 && cards[1].cnt == 1 {
		return 1
	}
	return 0
}

func compareHands(h1, h2 Hand, part2 bool) bool {
	score := map[byte]int{
		'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0,
	}
	if part2 {
		score = map[byte]int{
			'A': 12, 'K': 11, 'Q': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0,
		}
	}

	if handType(h1, part2) > handType(h2, part2) {
		return true
	} else if handType(h1, part2) == handType(h2, part2) {
		for k := range h1.hand {
			if score[h1.hand[k]] == score[h2.hand[k]] {
				continue
			}
			return score[h1.hand[k]] > score[h2.hand[k]]
		}
	}
	return false
}

// Solve part 1
func Solve1() {
	hands := []Hand{}

	for _, line := range lines {
		var h string
		var b int
		fmt.Sscanf(line, "%s %d", &h, &b)
		hands = append(hands, Hand{h, b})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], false)
	})

	for i := range hands {
		result += (len(hands) - i) * hands[i].bid
	}
}

// Solve part 2
func Solve2() {
	hands := []Hand{}

	for _, line := range lines {
		var h string
		var b int
		fmt.Sscanf(line, "%s %d", &h, &b)
		hands = append(hands, Hand{h, b})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], true)
	})

	for i := range hands {
		result += (len(hands) - i) * hands[i].bid
	}
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
