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
	inputFilename := getFileName(day, 1, example)
	inputFilepath := getFilePath(INPUT, inputFilename)

	input, err := getInput(inputFilepath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Did you make sure you saved the input to a file with the correct name?")
		os.Exit(1)
	}

	output := getOutput(day, star, input)
	if output == nil {
		fmt.Println("Please add dayXstarX function to getOutput() also implement it")
		os.Exit(1)
	}

	if !example {
		outputFilename := getFileName(day, star, false)
		outputFilepath := getFilePath(OUTPUT, outputFilename)
		saveOutput(output, outputFilepath, example)
	} else {
		exampleFilename := fmt.Sprintf("d%d_example.txt", day)
		exampleOutput, err := getInput(getFilePath(OUTPUT, exampleFilename))
		if err != nil {
			fmt.Printf("Please add the example output for day %d\n", day)
			os.Exit(1)
		}
		fmt.Printf("\nExpected Result:\n%s\n\n", string(exampleOutput))
	}

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
		sourceDir = "inputs"
	case OUTPUT:
		sourceDir = "outputs"
	}
	return pathToInputOutput + "/" + sourceDir + "/" + filename
}

func getInput(filepath string) ([]byte, error) {
	data, err := internals.GetData(filepath)
	if err != nil {
		return nil, fmt.Errorf("error getting input: %w", err)
	}
	return data, nil
}

func getOutput(day, star int, data []byte) []byte {
	switch day {
	case 1:
		if star == 1 {
			return day1star1(data)
		}
		return day1star2(data)
	case 2:
		if star == 1 {
			return day2star1(data)
		}
		return day2star2(data)
	case 3:
		if star == 1 {
			return day3star1(data)
		}
		return day3star2(data)
	case 4:
		if star == 1 {
			return day4star1(data)
		}
		return day4star2(data)
		/*
			case 5:
				if star == 1 {
					return day5star1(data)
				}
				return day5star2(data)
			case 6:
				if star == 1 {
					return day6star1(data)
				}
				return day6star2(data)
		*/
	}
	return nil
}

// func saveOutput(data, filename/path, example) calls either normally or uses saveExample???
func saveOutput(data []byte, filepath string, example bool) {
	if example {
		exampleFileData, err := internals.GetData(filepath)
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
