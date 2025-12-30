package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day3star1(data []byte) []byte {
	stringified := string(data)

	reg, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	if err != nil {
		// something went wrong
		fmt.Println("the regular expression didn't compile: %w", err)
		return data
	}

	sum := 0
	matches := reg.FindAllString(stringified, -1)
	for _, match := range matches {
		stringNumbers := strings.Trim(match, "mul()")
		splitStringNumbers := strings.Split(stringNumbers, ",")
		valueOne, _ := strconv.Atoi(splitStringNumbers[0])
		valueTwo, _ := strconv.Atoi(splitStringNumbers[1])
		sum += valueOne * valueTwo
	}

	return []byte(strconv.Itoa(sum))
}

func day3star2(data []byte) []byte {
	stringified := string(data)
	var dos strings.Builder

	//remove all sections covered by don't() until the next do()
	for {
		doString, remainingString, found := strings.Cut(stringified, "don't()")
		dos.WriteString(doString)
		if !found {
			break
		}

		_, stringified, found = strings.Cut(remainingString, "do()")
		if !found {
			break
		}
	}

	return day3star1([]byte(dos.String()))
}
