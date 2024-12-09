package main

import (
	"aoc/internal/reuse"
	"fmt"
)

func main() {
	filePath := "data/day1.txt"

	fileContent := reuse.ReadFile(filePath)
	listA, listB := splitSlices(fileContent)
	totalDistance := calculateDistance(listA, listB)
	totalScore := calculateScore(listA, listB)

	fmt.Printf("Distance: %d\n", totalDistance)
	fmt.Printf("Similarity Score: %d\n", totalScore)
}
