package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 19

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
	result = int(solve(lines, false))
}

// Solve part 2
func Solve2() {
	result = int(solve(lines, true))
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

type Range struct {
	start, end int64
}

type Rating map[string]int64

type Instruction struct {
	part           string
	quantity       int64
	sign           string
	targetWorkflow string
	follow         bool
	accept         bool
	reject         bool
}

type Workflow map[string][]Instruction

func executeWorkflow(ratings Rating, workflows Workflow, workflow string) bool {
	// sum := int64(0)
	// fmt.Println("Executing workflow:", workflow)
	instructions := workflows[workflow]
	for _, instruction := range instructions {
		_, ok := ratings[instruction.part]
		if ok {
			value := ratings[instruction.part]
			// If there is a instruction for this variable
			if instruction.sign == ">" && value > instruction.quantity {
				if instruction.follow {
					return executeWorkflow(ratings, workflows, instruction.targetWorkflow)
				} else if instruction.accept {
					return true
				} else if instruction.reject {
					return false
				}
			} else if instruction.sign == "<" && value < instruction.quantity {
				if instruction.follow {
					return executeWorkflow(ratings, workflows, instruction.targetWorkflow)
				} else if instruction.accept {
					return true
				} else if instruction.reject {
					return false
				}
			}

		} else if instruction.follow && instruction.part == "" {
			return executeWorkflow(ratings, workflows, instruction.targetWorkflow)
		} else if instruction.quantity == 0 && instruction.accept {
			return true
		} else if instruction.quantity == 0 && instruction.reject {
			//		fmt.Println("REJECT !")
			return false
		}

	}
	// Shouldn't be here, just being safe
	return false
}

func executeWorkflows(ratings Rating, workflows Workflow) int64 {

	// First execute the in workflow
	val := executeWorkflow(ratings, workflows, "in")
	// fmt.Println("Val:", val)
	sum := int64(0)
	if val {
		for _, rating := range ratings {
			sum += rating
		}
	}
	return sum
}

func doCombinations(workflows Workflow, workflow string, ranges []Range) int64 {
	fmt.Println(workflow, "Ranges: ", ranges)
	result := int64(0)
	instructions := workflows[workflow]
	for _, instruction := range instructions {
		rangeCopy := make([]Range, 4)
		copy(rangeCopy, ranges)
		indexVar := map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}[instruction.part]
		if instruction.sign == "<" {
			rangeCopy[indexVar].end = instruction.quantity - 1
			ranges[indexVar].start = instruction.quantity
			if instruction.follow {
				result += doCombinations(workflows, instruction.targetWorkflow, rangeCopy)
			} else if instruction.accept {
				result += doRangeProduct(rangeCopy)
			} else if instruction.reject {
				result += 0
			}
		} else if instruction.sign == ">" {
			rangeCopy[indexVar].start = instruction.quantity + 1
			ranges[indexVar].end = instruction.quantity
			if instruction.follow {
				result += doCombinations(workflows, instruction.targetWorkflow, rangeCopy)
			} else if instruction.accept {
				result += doRangeProduct(rangeCopy)
			} else if instruction.reject {
				result += 0
			}
		} else if instruction.accept {
			result += doRangeProduct(ranges)
		} else if instruction.reject {
			result += 0
		} else if instruction.follow {
			result += doCombinations(workflows, instruction.targetWorkflow, ranges)
		}
	}
	return result
}

func doRangeProduct(ranges []Range) int64 {
	result := int64(1)
	fmt.Println("Range product : Ranges", ranges)
	for _, r := range ranges {
		result *= r.end - r.start + 1
		fmt.Println("Result:", result)
	}
	fmt.Println("Result:", result)
	return result
}
func solve(input []string, part2 bool) int64 {
	decomposeLine := regexp.MustCompile(`^([a-z]+){(.*)}$`)
	decomposeInstructions := regexp.MustCompile(`([ARa-z]+)([<>]?)([0-9]+?):?([ARa-z]+?)$`)
	var workflows = make(Workflow)
	var sum int64
	inWorkflows := true
	for index, line := range input {
		if input[index] == "" {
			inWorkflows = false
			continue
		}
		if inWorkflows {
			var instructions []Instruction
			match := decomposeLine.FindStringSubmatch(line)
			// workflows.name = match[1]
			for _, instruction := range strings.Split(match[2], ",") {
				var targetInstruction Instruction
				matchI := decomposeInstructions.FindStringSubmatch(instruction)
				if len(matchI) > 0 {
					targetInstruction.part = matchI[1]
					targetInstruction.sign = matchI[2]
					targetInstruction.accept = matchI[4] == "A"
					targetInstruction.reject = matchI[4] == "R"
					if matchI[4] != "A" && matchI[4] != "R" {
						targetInstruction.targetWorkflow = matchI[4]
						targetInstruction.follow = true

					}
					num, _ := strconv.Atoi(matchI[3])
					targetInstruction.quantity = int64(num)
				} else if instruction == "A" || instruction == "R" {
					targetInstruction.accept = instruction == "A"
					targetInstruction.reject = instruction == "R"
				} else {
					targetInstruction.targetWorkflow = instruction
					targetInstruction.follow = true
				}

				instructions = append(instructions, targetInstruction)
				workflows[match[1]] = instructions

			}
			continue
		} else if !part2 {
			regexpParts := regexp.MustCompile(`{(.+)}`)
			rawParts := regexpParts.FindStringSubmatch(line)
			// fmt.Println(line, rawParts)
			parts := strings.Split(rawParts[1], ",")
			var ratings = Rating{}
			for _, part := range parts {
				//var rating Rating
				rawPart := strings.Split(part, "=")
				num, _ := strconv.Atoi(rawPart[1])
				ratings[rawPart[0]] = int64(num)
			}
			//			fmt.Println("Ratings:", ratings)
			val := executeWorkflows(ratings, workflows)
			//			fmt.Println(" Value:", val)

			sum += val
		}
	}
	if part2 {
		sum = doCombinations(workflows, "in", []Range{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}})
	}
	return sum
}
