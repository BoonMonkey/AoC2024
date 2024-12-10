package main

import (
	"aoc/internal/reuse"
	"log"
)

func main() {
	filePath := "data/day9.txt"
	fileContent := reuse.ReadFile(filePath)

	formattedString := formatString(fileContent)
	log.Print(formattedString)
	// splitString := strings.Split(formattedString, "")
	// log.Print(splitString)
	// structuredElements := structuredElementsRework(splitString)

	// calculatedCheckSum := calculateCheckSum(structuredElements)
	// log.Print(calculatedCheckSum)
}
