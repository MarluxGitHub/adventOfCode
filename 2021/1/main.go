package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inPutArray := readIntArrayFromFile("input.txt")
	sum := 0
	last := inPutArray[0]

	for i := 1; i < len(inPutArray); i++ {
		if last < inPutArray[i] {
			sum++
		}
		last = inPutArray[i]
	}

	fmt.Println(sum)

}

func readIntArrayFromFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// convert string to int
		i, _ := strconv.Atoi(scanner.Text())
		result = append(result, i)
	}
	return result
}