package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Flag parsing
	dayFlag := flag.Int("day", 0, "which day's puzzle solution code to run")
	exampleFlag := flag.Bool("example", false, "if set, runs example input instead of the real")
	flag.Parse()
	if *dayFlag == 0 {
		fmt.Println("Error: Day was not picked. Use -day followed by a number to pick which day to solve.")
		os.Exit(1)
	}

	fmt.Printf("Day picked: %d\n", *dayFlag)
	fmt.Printf("Run examples: %t\n", *exampleFlag)
}
