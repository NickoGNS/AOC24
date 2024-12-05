package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	col1, col2 := []int32{}, []int32{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		col1 = append(col1, int32(num1))
		col2 = append(col2, int32(num2))
	}

	sort.Slice(col1, func(i, j int) bool {
		return col1[i] < col1[j]
	})
	sort.Slice(col2, func(i, j int) bool {
		return col2[i] < col2[j]
	})

	var totalDistance int32 = 0
	similarityMap := make(map[int32]int32, len(col2))
	for i := range col1 {
		totalDistance += (func() int32 {
			if col1[i] < col2[i] {
				return col2[i] - col1[i]
			}
			return col1[i] - col2[i]
		})()
		similarityMap[col2[i]]++
	}

	var similarity int32 = 0
	for _, val := range col1 {
		if similarityMap[val] > 0 {
			similarity += val * similarityMap[val]
		}
	}

	fmt.Println(totalDistance)
	fmt.Println(similarity)
}
