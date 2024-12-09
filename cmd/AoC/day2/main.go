package main

import (
	"aoc/internal/reuse"
	"fmt"
)

func main() {
	filePath := "data/day2.txt"
	fileContent := reuse.ReadFile(filePath)
	safeIncrease := checkOrderIncreasing(fileContent)
	safeDecrease := checkOrderDecreasing(fileContent)

	var safeIncrementalIncrease [][]int
	var safeIncrementalDecrease [][]int
	for _, report := range safeIncrease {
		report := checkDiff(report)
		if report != nil {
			safeIncrementalIncrease = append(safeIncrementalIncrease, report)
		}
	}

	for _, report := range safeDecrease {
		report := checkDiff(report)
		if report != nil {
			safeIncrementalDecrease = append(safeIncrementalDecrease, report)
		}
	}

	fmt.Printf("Safe increase count: %d\n", len(safeIncrementalIncrease))
	fmt.Printf("Safe decrease count: %d\n", len(safeIncrementalDecrease))
	fmt.Printf("Safe reports: %d\n", len(safeIncrementalIncrease)+len(safeIncrementalDecrease))
}
