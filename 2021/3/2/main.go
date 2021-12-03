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
	return oxygen(inputs) * co2(inputs)
}

func oxygen(inputs []string) int64 {
	i := 0

	for len(inputs) > 1 {
		binaryNumbers := getBinaryNumbers(inputs)
		result := computeMostHit(binaryNumbers)

		oc := result[0][i]

		var inputs2 []string
		for j := 0; j < len(inputs); j++ {
			if(inputs[j][i] == oc) {
				inputs2 = append(inputs2, inputs[j])
			}
		}

		i++
		inputs = inputs2
	}

	return binaryToInt(inputs[0])
}

func co2(inputs []string) int64 {
	i := 0

	for len(inputs) > 1 {
		binaryNumbers := getBinaryNumbers(inputs)
		result := computeMostHit(binaryNumbers)

		oc := result[1][i]

		var inputs2 []string
		for j := 0; j < len(inputs); j++ {
			if(inputs[j][i] == oc) {
				inputs2 = append(inputs2, inputs[j])
			}
		}

		i++
		inputs = inputs2
	}

	return binaryToInt(inputs[0])
}


func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func computeMostHit(binaryNumbers []string) []string {
	result := make([]string, 2)
	for i := 0; i < len(binaryNumbers); i++ {
		c0, c1 := 0,0
		chars := []rune(binaryNumbers[i])
		for j:=0; j < len(chars); j++ {
			switch (chars[j]) {
				case '0': c0++
				case '1': c1++
			}
		}

		if(c1 >= c0) {
			result[0] += "1"
			result[1] += "0"
		} else {
			result[0] += "0"
			result[1] += "1"
		}
	}

	return result
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