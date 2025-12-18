package days

import (
	"adventofcode2024/internals"
	"fmt"
	"os"
)

const (
	INPUT = iota
	OUTPUT
)

func RunDay(day, star int, example bool) {
	filename := getFileName(day, star, example)
	filepath := getFilePath(INPUT, filename)
	url := getUrl(day, star)
	if url == "ERROR" {
		fmt.Println("ERROR: add the url for the new day!!!")
		os.Exit(1)
	}

	data, err := getInput(filepath, url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := day1star1(data)
	outputFilepath := getFilePath(OUTPUT, filename)
	saveOutput(output, outputFilepath, example)
	fmt.Println("Output: " + string(output))
}

func getFileName(day, star int, example bool) string {
	if example {
		return fmt.Sprintf("d%ds%d_example.txt", day, star)
	}
	return fmt.Sprintf("d%ds%d.txt", day, star)
}

func getFilePath(source int, filename string) string {
	var sourceDir string
	switch source {
	case INPUT:
		sourceDir = "input"
	case OUTPUT:
		sourceDir = "output"
	}
	return pathToInputOutput + sourceDir + filename
}

func getInput(filepath, url string) ([]byte, error) {
	data, err := internals.GetData(url, filepath)
	if err != nil {
		return nil, fmt.Errorf("error getting input: %w", err)
	}
	return data, nil
}

// func saveOutput(data, filename/path, example) calls either normally or uses saveExample???
func saveOutput(data []byte, filepath string, example bool) {
	if example {
		exampleFileData, err := internals.GetData("", filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data = append(exampleFileData, data...)
	}

	err := internals.SaveToFile(data, filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
