package main

import "fmt"

func main() {
	filePath := "list.txt"

	fileContent := readFile(filePath)
	listA, listB := splitSlices(fileContent)
	totalDistance := calculateDistance(listA, listB)
	totalScore := calculateScore(listA, listB)

	fmt.Printf("Distance: %d\n", totalDistance)
	fmt.Printf("Similarity Score: %d\n", totalScore)
}
