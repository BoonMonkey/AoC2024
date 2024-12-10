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
	splitString := strings.Split(formattedString, "_split")
	if splitString[len(splitString)-1] == "" {
		splitString = splitString[:len(splitString)-1]
	}
	structuredElements := structuredElementsRework(splitString)

	calculatedCheckSum := calculateCheckSum(structuredElements)
	log.Print(calculatedCheckSum)
}
