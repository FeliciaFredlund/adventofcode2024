package days

import (
	"strconv"
	"strings"
)

func day4star1(data []byte) []byte {
	stringified := string(data)
	matrix := [][]string{}

	for line := range strings.Lines(stringified) {
		matrix = append(matrix, strings.Split(strings.TrimSpace(line), ""))
	}

	count := 0
	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			if matrix[i][j] == "X" {
				count += lookForMAS(matrix, i, j)
			}
		}
	}

	return []byte(strconv.Itoa(count))
}

func day4star2(data []byte) []byte {
	stringified := string(data)
	matrix := [][]string{}

	for line := range strings.Lines(stringified) {
		matrix = append(matrix, strings.Split(strings.TrimSpace(line), ""))
	}

	count := 0
	for i := range len(matrix) {
		if i == 0 || i == len(matrix)-1 {
			continue
		}

		for j := range len(matrix[0]) {
			if j == 0 || j == len(matrix[0])-1 {
				continue
			}

			if matrix[i][j] == "A" {
				if lookForMS(matrix, i, j) {
					count++
				}
			}
		}
	}

	return []byte(strconv.Itoa(count))
}

func lookForMAS(matrix [][]string, xi, xj int) int {
	countMAS := 0

	// north direction
	if xi-3 >= 0 {
		// north
		if matrix[xi-1][xj] == "M" {
			if matrix[xi-2][xj] == "A" {
				if matrix[xi-3][xj] == "S" {
					countMAS++
				}
			}
		}

		// northeast
		if xj+3 < len(matrix[0]) {
			if matrix[xi-1][xj+1] == "M" {
				if matrix[xi-2][xj+2] == "A" {
					if matrix[xi-3][xj+3] == "S" {
						countMAS++
					}
				}
			}
		}

		// northwest
		if xj-3 >= 0 {
			if matrix[xi-1][xj-1] == "M" {
				if matrix[xi-2][xj-2] == "A" {
					if matrix[xi-3][xj-3] == "S" {
						countMAS++
					}
				}
			}
		}
	}

	// east
	if xj+3 < len(matrix[0]) {
		if matrix[xi][xj+1] == "M" {
			if matrix[xi][xj+2] == "A" {
				if matrix[xi][xj+3] == "S" {
					countMAS++
				}
			}
		}
	}

	// west
	if xj-3 >= 0 {
		if matrix[xi][xj-1] == "M" {
			if matrix[xi][xj-2] == "A" {
				if matrix[xi][xj-3] == "S" {
					countMAS++
				}
			}
		}
	}

	// south direction
	if xi+3 < len(matrix) {
		// south
		if matrix[xi+1][xj] == "M" {
			if matrix[xi+2][xj] == "A" {
				if matrix[xi+3][xj] == "S" {
					countMAS++
				}
			}
		}

		// southeast
		if xj+3 < len(matrix[0]) {
			if matrix[xi+1][xj+1] == "M" {
				if matrix[xi+2][xj+2] == "A" {
					if matrix[xi+3][xj+3] == "S" {
						countMAS++
					}
				}
			}
		}

		// southwest
		if xj-3 >= 0 {
			if matrix[xi+1][xj-1] == "M" {
				if matrix[xi+2][xj-2] == "A" {
					if matrix[xi+3][xj-3] == "S" {
						countMAS++
					}
				}
			}
		}
	}

	return countMAS
}

func lookForMS(matrix [][]string, ai, aj int) bool {
	if ai-1 < 0 || ai+1 == len(matrix) || aj-1 < 0 || aj+1 == len(matrix[0]) {
		return false
	}

	if matrix[ai-1][aj-1] == "M" && matrix[ai+1][aj+1] == "S" {
		if matrix[ai-1][aj+1] == "M" && matrix[ai+1][aj-1] == "S" {
			return true
		}
		if matrix[ai-1][aj+1] == "S" && matrix[ai+1][aj-1] == "M" {
			return true
		}
	}
	if matrix[ai-1][aj-1] == "S" && matrix[ai+1][aj+1] == "M" {
		if matrix[ai-1][aj+1] == "M" && matrix[ai+1][aj-1] == "S" {
			return true
		}
		if matrix[ai-1][aj+1] == "S" && matrix[ai+1][aj-1] == "M" {
			return true
		}
	}

	return false
}
