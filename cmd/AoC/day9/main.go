package main

import (
	"aoc/internal/reuse"
	"log"
	"strings"
)

func main() {
	filePath := "data/day9.txt"
	fileContent := reuse.ReadFile(filePath)

	formattedString := formatString(fileContent)
	splitString := strings.Split(formattedString, "_split") // J
	if splitString[len(splitString)-1] == "" {              // A
		splitString = splitString[:len(splitString)-1] // N
	} // K
	structuredElements := structuredElementsRework(splitString)

	calculatedCheckSum := calculateCheckSum(structuredElements)
	log.Print(calculatedCheckSum)
}
