package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/echojc/aocutil"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }
func printf(f string)  { fmt.Fprintf(writer, f) }

var lines []string
var result int
var year = 2023
var day = 6

type Race struct {
	Time, Distance int
}

var Races []Race = make([]Race, 0)
var RealRace Race = Race{}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	readInput()
	parseRaces()
	parseRealRace()

	result = 0
	Solve1()
	println("1:" + strconv.Itoa(result))

	result = 0
	Solve2()
	println("2:" + strconv.Itoa(result))
}

// Solve part 1
func Solve1() {
	for _, race := range Races {
		count := 0

		for buttonHoldTime := 0; buttonHoldTime <= race.Time; buttonHoldTime++ {
			distance := 0
			for time := buttonHoldTime + 1; time <= race.Time; time++ {
				distance += buttonHoldTime
				if distance > race.Distance {
					count++
					break
				}
			}
		}

		if result == 0 {
			result = count
		} else {
			result *= count
		}
	}
}

// Solve part 2
func Solve2() {
	result = solve(RealRace.Time, RealRace.Distance)
}

func solve(time, distance int) int {
	minimum := (float64(RealRace.Time) - math.Sqrt(math.Pow(float64(RealRace.Time), 2)-4*float64(RealRace.Distance))/2)
	maximum := (float64(RealRace.Time) + math.Sqrt(math.Pow(float64(RealRace.Time), 2)-4*float64(RealRace.Distance))/2)

	minHoldTime := math.Floor(minimum + 1)
	maxHoldTime := math.Ceil(maximum - 1)

	return int(maxHoldTime - minHoldTime + 1)
}

func parseRaces() {
	re := regexp.MustCompile(`(\d+)`)

	matches := re.FindAllStringSubmatch(lines[0], -1)

	if len(matches) > 0 {
		for _, match := range matches {
			race := Race{}

			time, err := strconv.Atoi(match[1])

			if err != nil {
				log.Fatal(err)
			}

			race.Time = time
			Races = append(Races, race)
		}
	}

	matches = re.FindAllStringSubmatch(lines[1], -1)

	if len(matches) > 0 {
		for i, match := range matches {
			distance, err := strconv.Atoi(match[1])

			if err != nil {
				log.Fatal(err)
			}

			Races[i].Distance = distance
		}
	}
}

func parseRealRace() {
	time := ""
	distance := ""

	for _, race := range Races {
		time += strconv.Itoa(race.Time)
		distance += strconv.Itoa(race.Distance)
	}

	timeInt, err := strconv.Atoi(time)

	if err != nil {
		log.Fatal(err)
	}

	distanceInt, err := strconv.Atoi(distance)

	if err != nil {
		log.Fatal(err)
	}

	RealRace.Time = timeInt
	RealRace.Distance = distanceInt
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
