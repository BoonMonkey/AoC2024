package main

import (
	"log"
	"strconv"
	"strings"
)

func checkOrderIncreasing(fileContent []string) [][]int {
	var safeIncrease [][]int
	for _, line := range fileContent {
		var safetyOrderIncreasing []int
		var safetyCount int
		splitLine := strings.Split(line, " ")

		for i := 0; i < len(splitLine); i++ {
			currentItem, err := strconv.Atoi(splitLine[i])
			if err != nil {
				log.Fatal(err)
			}

			if len(safetyOrderIncreasing) == 0 {
				safetyOrderIncreasing = append(safetyOrderIncreasing, currentItem)
				safetyCount = 0
			}

			if currentItem > safetyOrderIncreasing[safetyCount] {
				safetyOrderIncreasing = append(safetyOrderIncreasing, currentItem)
				safetyCount++
			}
		}

		if len(safetyOrderIncreasing) == len(splitLine) {
			// log.Print("Safe")
			safeIncrease = append(safeIncrease, safetyOrderIncreasing)
		} else {
			// log.Print("Unsafe")
		}
	}

	return safeIncrease
}

func checkOrderDecreasing(fileContent []string) [][]int {
	var safeDecrease [][]int
	for _, line := range fileContent {
		var safetyOrderDecreasing []int
		var safetyCount int
		splitLine := strings.Split(line, " ")

		for i := 0; i < len(splitLine); i++ {
			currentItem, err := strconv.Atoi(splitLine[i])
			if err != nil {
				log.Fatal(err)
			}

			if len(safetyOrderDecreasing) == 0 {
				safetyOrderDecreasing = append(safetyOrderDecreasing, currentItem)
				safetyCount = 0
			}

			if currentItem < safetyOrderDecreasing[safetyCount] {
				safetyOrderDecreasing = append(safetyOrderDecreasing, currentItem)
				safetyCount++
			}
		}

		if len(safetyOrderDecreasing) == len(splitLine) {
			safeDecrease = append(safeDecrease, safetyOrderDecreasing)
		} else {
			// log.Print("Unsafe")
		}
	}

	return safeDecrease
}
