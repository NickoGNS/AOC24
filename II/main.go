package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func isValidSequence(levels []int, hasProblemDampener bool) bool {
	positions := []int{-1}
	if hasProblemDampener {
		for i := range levels {
			positions = append(positions, i)
		}
	}

	for _, skipPos := range positions {
		increasing := true
		decreasing := true
		prev := -1

		for i := range levels {
			if i == skipPos {
				continue
			}

			if prev == -1 {
				prev = i
				continue
			}

			curr := levels[i]
			prevVal := levels[prev]

			diff := abs(curr - prevVal)
			if diff < 1 || diff > 3 {
				increasing = false
				decreasing = false
				break
			}

			if curr <= prevVal {
				increasing = false
			}
			if curr >= prevVal {
				decreasing = false
			}

			prev = i
		}

		if increasing || decreasing {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	safeRepCount := 0
	safeRepCountPD := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		report := scanner.Text()
		strLevels := strings.Fields(report)
		levels := make([]int, 0, len(strLevels))

		valid := true
        for _, s := range strLevels {
            n, err := strconv.Atoi(s)
            if err != nil {
                valid = false
                break
            }
            levels = append(levels, n)
        }

        if !valid {
            continue
        }

		if isValidSequence(levels, false) {
            safeRepCount++
        }
		if isValidSequence(levels, true) {
			safeRepCountPD++
		}
	}

	fmt.Println(safeRepCount)
	fmt.Println(safeRepCountPD)
}
