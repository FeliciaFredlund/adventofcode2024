package days

import (
	"slices"
	"strconv"
	"strings"
)

// day 1: pair up values in two lists, sorted in ascending order
// and get the difference between the numbers
// sum those differences together

// func parseInput into two slices
// func getDifference between values

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
	// IMPLEMENT
	return data
}
