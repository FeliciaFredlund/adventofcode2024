package main

import (
	"adventofcode2024/days"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Flag parsing
	dayFlag := flag.Int("day", 0, "which day's puzzle solution code to run")
	starFlag := flag.Int("star", 1, "which star of the day to run, defaults to first star")
	exampleFlag := flag.Bool("example", false, "if set, runs example input instead of the real")
	flag.Parse()
	if *dayFlag == 0 {
		fmt.Println("Error: Day was not picked. Use -day followed by a number to pick which day to solve.")
		os.Exit(1)
	}

	fmt.Printf("Day picked: %d\n", *dayFlag)
	fmt.Printf("Star picked: %d\n", *starFlag)
	fmt.Printf("Run examples: %t\n\n", *exampleFlag)
	days.RunDay(*dayFlag, *starFlag, *exampleFlag)
}
