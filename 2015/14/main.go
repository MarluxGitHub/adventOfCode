package main

import (
	"MarluxGitHub/adventOfCode/pkg/math"
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
var year = 2015
var day = 14

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
	time := 2503

	rendeerDistance := make(map[string]int)

	// example String: Vixen can fly 18 km/s for 5 seconds, but then must rest for 84 seconds.
	for _, line := range lines {
		var name string
		var speed, flyTime, restTime int

		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&name, &speed, &flyTime, &restTime)
		if err != nil {
			log.Fatal(err)
		}

		cycleTime := flyTime + restTime
		fullCycles := time / cycleTime
		remainingTime := time % cycleTime

		distance := fullCycles*speed*flyTime + math.Min(remainingTime, flyTime)*speed
		rendeerDistance[name] = distance
	}

	maxDistance := 0
	for _, distance := range rendeerDistance {
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	result = maxDistance
}

// Solve part 2
func Solve2() {
	time := 2503

	type reendeer struct {
		speed, flyTime, restTime int
		points                   int
	}

	reendeers := make(map[string]reendeer)

	for _, line := range lines {
		var name string
		var speed, flyTime, restTime int

		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&name, &speed, &flyTime, &restTime)
		if err != nil {
			log.Fatal(err)
		}

		reendeers[name] = reendeer{speed, flyTime, restTime, 0}
	}

	for i := 1; i <= time; i++ {
		reendeerDistance := make(map[string]int)

		for name, reendeer := range reendeers {
			cycleTime := reendeer.flyTime + reendeer.restTime
			fullCycles := i / cycleTime
			remainingTime := i % cycleTime

			distance := fullCycles*reendeer.speed*reendeer.flyTime + math.Min(remainingTime, reendeer.flyTime)*reendeer.speed
			reendeerDistance[name] = distance
		}

		maxDistance := 0
		for _, distance := range reendeerDistance {
			if distance > maxDistance {
				maxDistance = distance
			}
		}

		for name, distance := range reendeerDistance {
			if distance == maxDistance {
				temp := reendeers[name]
				temp.points++
				reendeers[name] = temp
			}
		}

	}

	maxPoints := 0
	for _, reendeer := range reendeers {
		if reendeer.points > maxPoints {
			maxPoints = reendeer.points
		}
	}
	result = maxPoints

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
