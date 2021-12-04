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

type BoardEntry struct {
	value string
	marked bool
}

var bingoInput []string
type BingoBoard [5][5]BoardEntry
var bingoBoards []BingoBoard


func (b BingoBoard) fillBingoBoard(input []string) BingoBoard {
	for i := 0; i < 5; i++ {
		line := strings.Split(standardizeSpaces(input[i]), " ")
		for j := 0; j < len(line); j++ {
			b[i][j] = BoardEntry{line[j], false}
		}
	}

	return b
}

func (b BingoBoard) print() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if(!b[i][j].marked) {
				printf(b[i][j].value)
			} else {
				printf("_")
			}
			printf(" ")
		}
		println("")
	}
}

func (b BingoBoard) findEntryInBoard(number string) BingoBoard {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// println(number + " " + b[i][j].value)
			if b[i][j].value == number {
				b[i][j].marked = true
			}
		}
	}

	return b
}

func (b BingoBoard) sumBoard(input string) int {
	sum := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
		if !b[i][j].marked {
			number, _ := strconv.Atoi(b[i][j].value)
			sum += number
			}
		}
	}

	println(input + " " + strconv.Itoa(sum))

	number, _ := strconv.Atoi(input)
	sum = sum * number

	b.print()

	return sum
}

func (b BingoBoard) hasBingo() bool {
	for i := 0; i < 5; i++ {
		if b.checkRow(i) {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		if b.checkColumn(i) {
			return true
		}
	}

	// if b.checkDiagonal() {
	// 	return true
	// }

	return false
}

func (b BingoBoard) checkRow(row int) bool {
	sum := 0
	for i := 0; i < 5; i++ {
		if b[row][i].marked {
			sum++
		}
	}

	if(sum == 5) {
		return true
	} else {
		return false
	}
}

func (b BingoBoard) checkColumn(column int) bool {
	sum := 0

	for i := 0; i < 5; i++ {
		if b[i][column].marked {
			sum++
		}
	}

	if(sum == 5) {
		return true
	} else {
		return false
	}
}

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readInput()

  var ignoredBoards []int

  out:
  for i := 0; i < len(bingoInput); i++ {
	println(bingoInput[i])
	println("")

  	for j := 0; j < len(bingoBoards); j++ {
  		bingoBoards[j] = bingoBoards[j].findEntryInBoard(bingoInput[i])
		if(bingoBoards[j].hasBingo()) {
			// array contains j
			found := false
			for k := 0; k < len(ignoredBoards); k++ {
				if(ignoredBoards[k] == j) {
					found = true
				}
			}

			if(!found) {
				ignoredBoards = append(ignoredBoards, j)
			}
		}
  	}
	  if(len(ignoredBoards) == len(bingoBoards)) {
		println("Bingo")
		println(strconv.Itoa(bingoBoards[ignoredBoards[len(ignoredBoards)-1]].sumBoard(bingoInput[i])))
		break out
	  }
  }

}

func removeBoardIndex(index int) {
	for i := index; i < len(bingoBoards) - 1; i++ {
		bingoBoards[i] = bingoBoards[i+1]
	}

	bingoBoards = bingoBoards[:len(bingoBoards) - 1]
}

func readInput() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
	log.Fatal(err)
	}

	lines, err := i.Strings(2021, 4)
	if err != nil {
	log.Fatal(err)
	}

	bingoInput = strings.Split(lines[0], ",")

	for i := 1; i < len(lines); i++ {
		if(lines[i] == "") {
			continue
		}

		board := BingoBoard{}
		board = board.fillBingoBoard(lines[i:i+5])
		bingoBoards = append(bingoBoards, board)
		i += 5
	}
}
