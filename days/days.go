package days

import (
	"adventofcode2024/internals"
	"fmt"
)

const (
	input = iota
	output
)

func getFileName(day int, star int, example bool) string {
	if example {
		return fmt.Sprintf("d%ds%d_example.txt", day, star)
	}
	return fmt.Sprintf("d%ds%d.txt", day, star)
}

func getFilePath(source int, filename string) string {
	var sourceDir string
	switch source {
	case input:
		sourceDir = "input"
	case output:
		sourceDir = "output"
	}
	return pathToInputOutput + sourceDir + filename
}

func getInput(filepath, url string) []byte {
	data, err := internals.GetData(url, filepath)
	if err != nil {
		//DO STUFF
	}

	return data
}

// func saveOutput(data, filename/path, example) calls either normally or uses saveExample???
func saveOutput(data []byte, filepath string, example bool) {
	if example {
		// do things specific for example, like read in the input file and append the data slice to it
	}
	// call on io SaveToFile(data, filepath)
}
