package main

import (
	"aoc/internal/reuse"
	"log"
)

func main() {
	filePath := "data/day2.txt"

	fileContent := reuse.ReadFile(filePath)
	log.Print(fileContent[1])
}
