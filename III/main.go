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

	fileStr := string(file)
	mulSum := 0
	isEnabled := true

	combined := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	for _, match := range combined.FindAllStringSubmatchIndex(fileStr, -1) {
		text := fileStr[match[0]:match[1]]

		switch {
		case text == "don't()":
			isEnabled = false
		case text == "do()":
			isEnabled = true
		case isEnabled && text[:4] == "mul(":
			num1, _ := strconv.Atoi(fileStr[match[2]:match[3]])
			num2, _ := strconv.Atoi(fileStr[match[4]:match[5]])
			mulSum += num1 * num2
		}
	}

	fmt.Println(mulSum)
}
