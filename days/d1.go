package days

import (
	"slices"
	"strconv"
	"strings"
)

func day1star1(data []byte) []byte {
	stringified := string(data)
	numbers := strings.Fields(stringified)
	leftList := []int{}
	rightList := []int{}
	for i, number := range numbers {
		value, _ := strconv.Atoi(number)
		if i%2 == 0 {
			leftList = append(leftList, value)
		} else {
			rightList = append(rightList, value)
		}
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDifference := 0
	for i, left := range leftList {
		difference := left - rightList[i]
		if difference < 0 {
			difference = -difference
		}
		totalDifference += difference
	}

	return []byte(strconv.Itoa(totalDifference))
}

func day1star2(data []byte) []byte {
	stringified := string(data)
	numbers := strings.Fields(stringified)
	leftList := []string{}
	rightList := map[string]int{}
	for i, number := range numbers {
		if i%2 == 0 {
			leftList = append(leftList, number)
		} else {
			rightList[number]++
		}
	}

	similarityScore := 0
	for _, item := range leftList {
		value, _ := strconv.Atoi(item)
		similarityScore += (value * rightList[item])
	}

	return []byte(strconv.Itoa(similarityScore))
}
