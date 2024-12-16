package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	normalized := strings.ReplaceAll(string(file), "\r\n", "\n")
	fileParts := strings.Split(normalized, "\n\n")
	fmt.Println(fileParts[1])

	mapping := make(map[string]map[string]struct{})
	// afterMap := make(map[string]string)
	for _, line := range strings.Split(fileParts[0], "\n") {
		lineParts := strings.Split(line, "|")
		fmt.Println(lineParts[0])

		key := strings.TrimSpace(lineParts[0])
		value := strings.TrimSpace(lineParts[1])

		if _, exists := mapping[key]; !exists {
			mapping[key] = make(map[string]struct{})
		}

		mapping[key][value] = struct{}{}
	}

	fmt.Println(mapping["83"])

	// isEnabled := true
	//
	// combined := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	//
	// for _, match := range combined.FindAllStringSubmatchIndex(fileStr, -1) {
	// text := fileStr[match[0]:match[1]]
	//
	// switch {
	// case text == "don't()":
	// isEnabled = false
	// case text == "do()":
	// isEnabled = true
	// case isEnabled && text[:4] == "mul(":
	// num1, _ := strconv.Atoi(fileStr[match[2]:match[3]])
	// num2, _ := strconv.Atoi(fileStr[match[4]:match[5]])
	// mulSum += num1 * num2
	// }
	// }

	// fmt.Println(mulSum)
}
