package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	direction string
	speed int
}

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

var directions [] Submarine

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  readFile("../input.txt")
  response := computeDirections()
  printf("%d\n", response)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		speed, _ := strconv.Atoi(tokens[1])
		directions = append(directions, Submarine{tokens[0], speed})
	}
}

func computeDirections() int {
	depth, forward, aim := 0, 0, 0

	for _, direction := range directions {
		switch direction.direction {
			case "forward":
				forward += direction.speed
				depth += aim*direction.speed
			case "down": aim += direction.speed
			case "up": aim -= direction.speed
		}
	}

	return depth * forward
}