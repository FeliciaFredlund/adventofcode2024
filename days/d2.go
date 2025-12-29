package days

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func day2star1(data []byte) []byte {
	stringified := string(data)
	safe := 0
	for line := range strings.Lines(stringified) {
		stringNumbers := strings.Fields(line)

		isSafe, _ := lineIsSafe(stringNumbers)
		if isSafe {
			safe++
		}
	}

	return []byte(strconv.Itoa(safe))
}

func day2star2(data []byte) []byte {
	// CLEAN THIS SHIT UP, THE CASE I MISSED WAS A FIRST INCREASE/DECREASE THAT WAS THE OUTLIER
	// AND THE REST OF THE REPORT WAS THE OTHER
	// BUT I DON'T THINK I NEED OBEBADREPORTFOUND BUT DELETING IT MADE THINGS BREAK

	stringified := string(data)
	safe := 0
	for line := range strings.Lines(stringified) {
		stringNumbers := strings.Fields(line)
		oneBadReportFound := false
		isSafe, indexOfUnsafe := lineIsSafe(stringNumbers)
		if isSafe {
			safe++
		} else if !oneBadReportFound {
			oneBadReportFound = true
			if isSafe, _ := lineIsSafe(slices.Delete(slices.Clone(stringNumbers), indexOfUnsafe, indexOfUnsafe+1)); isSafe {
				fmt.Println("removing first")
				safe++
			} else if isSafe, _ := lineIsSafe(slices.Delete(slices.Clone(stringNumbers), indexOfUnsafe+1, indexOfUnsafe+2)); isSafe {
				fmt.Println("removing second")
				safe++
			} else if indexOfUnsafe > 0 {
				if isSafe, _ := lineIsSafe(slices.Delete(slices.Clone(stringNumbers), indexOfUnsafe-1, indexOfUnsafe)); isSafe {
					fmt.Println("removing second")
					safe++
				}
			}
		}
	}

	return []byte(strconv.Itoa(safe))
}

func lineIsSafe(line []string) (bool, int) {
	increasing := false
	decreasing := false
	indexOfUnsafe := -1
	for i := range len(line) {
		if i+1 == len(line) {
			return true, -1
		}
		currentValue, _ := strconv.Atoi(line[i])
		nextValue, _ := strconv.Atoi(line[i+1])

		difference := math.Abs(float64(currentValue - nextValue))

		if difference == 0 || difference > 3 {
			indexOfUnsafe = i
			break
		}

		if currentValue < nextValue {
			increasing = true
		} else {
			decreasing = true
		}

		// if a line has both increasing and decreasing or two values are same, break out, this report ain't safe
		if increasing && decreasing {
			indexOfUnsafe = i
			break
		}
	}
	return false, indexOfUnsafe
}
