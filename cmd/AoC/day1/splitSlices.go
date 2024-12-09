package main

import (
	"log"
	"strconv"
	"strings"
)

func splitSlices(fileContent []string) ([]int, []int) {
	var listA []int
	var listB []int

	for _, line := range fileContent {
		splitLines := strings.Split(line, "   ")
		itemA, err := strconv.Atoi(splitLines[0])
		if err != nil {
			log.Fatal(err)
		}

		itemB, err := strconv.Atoi(splitLines[1])
		if err != nil {
			log.Fatal(err)
		}
		listA = append(listA, itemA)
		listB = append(listB, itemB)

	}

	return listA, listB
}
