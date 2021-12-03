package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func println(f string) { fmt.Fprintln(writer, f) }

func main() {
  // STDOUT MUST BE FLUSHED MANUALLY!!!
  defer writer.Flush()
  inputs := readFile("../input.txt")
  response := transform(inputs)
  println(strconv.Itoa(int(response)))
}

func readFile(filename string) []string {
	var inputs []string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}

func transform(inputs []string) int64 {
	result := make([]string, 2)
	binaryNumbers := getBinaryNumbers(inputs)

	for i := 0; i < len(binaryNumbers); i++ {
		c0, c1 := 0,0
		chars := []rune(binaryNumbers[i])
		for j:=0; j < len(chars); j++ {
			switch (chars[j]) {
				case '0': c0++
				case '1': c1++
			}
		}

		if(c1 > c0) {
			result[0] += "1"
			result[1] += "0"
		} else {
			result[0] += "0"
			result[1] += "1"
		}
	}

	return binaryToInt(result[0]) * binaryToInt(result[1])
}

func getBinaryNumbers(inputs []string) []string {
	var binaryNumbers []string
	binaryNumbers = make([]string, len(inputs[0]))

	for i := 0; i < len(inputs); i++ {
		chars := []rune(inputs[i])
		for j:=0; j < len(chars); j++ {
			binaryNumbers[j] += string(chars[j])
		}
	}

	return binaryNumbers
}


func binaryToInt(binary string) int64 {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return i
}