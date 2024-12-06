package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	mulSum := 0

	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := reg.FindAllStringSubmatch(string(file), -1)

	for _, match := range matches {
		if len(match) == 3 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			mulSum += (num1 * num2)
		}
	}

	fmt.Println(mulSum)
	// fmt.Println(mulSum)
}
