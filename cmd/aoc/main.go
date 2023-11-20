package main

import (
	"MarluxGitHub/adventOfCode/pkg/loader"
	"flag"
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
