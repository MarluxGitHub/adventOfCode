package main

import (
	"flag"
	"marlux/aoc/internal/loader"
	"os"
)

func main() {
	year := flag.Int("y", 0, "year")
	day := flag.Int("d", 0, "day")

	flag.Parse()

	if *year == 0 || *day == 0 {
		flag.Usage()
		os.Exit(1)
	}

	loader.Generate(*year, *day)
}