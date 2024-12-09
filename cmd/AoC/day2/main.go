package main

import (
	"aoc/internal/reuse"
	"fmt"
)

func main() {
	filePath := "data/day2.txt"
	fileContent := reuse.ReadFile(filePath)
	safeIncrease, _ := checkOrderIncreasing(fileContent)
	safeDecrease, _ := checkOrderDecreasing(fileContent)

	// tests
	test := checkOrderIncreasingRework(fileContent)
	test2 := checkOrderDecreasingRework(fileContent)

	var safeIncrementalIncrease [][]int
	var safeIncrementalDecrease [][]int
	for _, report := range safeIncrease {
		report := checkDiff(report)
		checkPotentials(report)
		if report != nil {
			safeIncrementalIncrease = append(safeIncrementalIncrease, report)
		}
	}

	for _, report := range safeDecrease {
		report := checkDiff(report)
		checkPotentials(report)
		if report != nil {
			safeIncrementalDecrease = append(safeIncrementalDecrease, report)
		}
	}

	fmt.Printf("Safe increase count: %d\n", len(safeIncrementalIncrease))
	fmt.Printf("Safe decrease count: %d\n", len(safeIncrementalDecrease))
	fmt.Printf("Rework safe increase: %d\n", len(test))
	fmt.Printf("Rework safe decrease: %d\n", len(test2))

	fmt.Printf("Rework total safe: %d\n", len(test)+len(test2))

	fmt.Printf("Safe reports: %d\n",
		len(safeIncrementalIncrease)+
			len(safeIncrementalDecrease))
}
