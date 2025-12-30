package days

import (
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
	stringified := string(data)

	safe := 0
	for line := range strings.Lines(stringified) {
		stringNumbers := strings.Fields(line)
		isSafe, indexOfUnsafe := lineIsSafe(stringNumbers)
		if isSafe {
			safe++
			continue
		}
		if indexOfUnsafe == -1 {
			//something went wrong, assume report is unsafe no matter what
			continue
		}

		// checking if removing one element would make it a safe report
		// remove element at left side of unsafe, remove element at right side of unsafe,
		// if index is 1, check if the unsafe is actually between the first and second element and is inc/dec mismatch
		if isSafe, _ := lineIsSafe(slices.Delete(slices.Clone(stringNumbers), indexOfUnsafe, indexOfUnsafe+1)); isSafe {
			safe++
		} else if isSafe, _ := lineIsSafe(slices.Delete(slices.Clone(stringNumbers), indexOfUnsafe+1, indexOfUnsafe+2)); isSafe {
			safe++
		} else if indexOfUnsafe == 1 {
			if isSafe, _ := lineIsSafe(slices.Delete(slices.Clone(stringNumbers), indexOfUnsafe-1, indexOfUnsafe)); isSafe {
				safe++
			}
		}
	}

	return []byte(strconv.Itoa(safe))
}

func lineIsSafe(line []string) (bool, int) {
	increasing := false
	decreasing := false
	for i := range len(line) {
		if i+1 == len(line) {
			return true, -1
		}
		currentValue, _ := strconv.Atoi(line[i])
		nextValue, _ := strconv.Atoi(line[i+1])

		difference := math.Abs(float64(currentValue - nextValue))

		if difference == 0 || difference > 3 {
			return false, i
		}

		if currentValue < nextValue {
			increasing = true
		} else {
			decreasing = true
		}

		// if a line has both increasing and decreasing, this report ain't safe
		if increasing && decreasing {
			return false, i
		}
	}
	return false, -1
}
