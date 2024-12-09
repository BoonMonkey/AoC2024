package main

import (
	"log"
	"strconv"
	"strings"
)

func checkOrderIncreasing(fileContent []string) ([][]int, [][]int) {
	var safeIncrease [][]int
	var unsafeIncrease [][]int
	for _, line := range fileContent {
		var safetyOrderIncreasing []int
		var unsafeOrderIncreasing []int
		var safetyCount int

		splitLine := strings.Split(line, " ")

		for i := 0; i < len(splitLine); i++ {
			currentItem, err := strconv.Atoi(splitLine[i])
			if err != nil {
				log.Fatal(err)
			}

			if len(safetyOrderIncreasing) == 0 {
				safetyOrderIncreasing = append(safetyOrderIncreasing, currentItem)
				unsafeOrderIncreasing = append(unsafeOrderIncreasing, currentItem)
				safetyCount = 0
			}

			if currentItem > safetyOrderIncreasing[safetyCount] {
				safetyOrderIncreasing = append(safetyOrderIncreasing, currentItem)
				safetyCount++
			} else {
				unsafeOrderIncreasing = append(unsafeOrderIncreasing, currentItem)
			}
		}

		if len(safetyOrderIncreasing) == len(splitLine) {
			// log.Print("Safe")
			safeIncrease = append(safeIncrease, safetyOrderIncreasing)
		}

		if len(unsafeOrderIncreasing) == len(splitLine) {
			unsafeIncrease = append(unsafeIncrease, unsafeOrderIncreasing)
		}
	}

	return safeIncrease, unsafeIncrease
}

func checkOrderDecreasing(fileContent []string) ([][]int, [][]int) {
	var safeDecrease [][]int
	var unsafeDecrease [][]int
	for _, line := range fileContent {
		var safetyOrderDecreasing []int
		var unsafeOrderDecreasing []int
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
			} else if len(safetyOrderDecreasing) == 0 {
				unsafeOrderDecreasing = append(unsafeOrderDecreasing, currentItem)
			}

			if currentItem < safetyOrderDecreasing[safetyCount] {
				safetyOrderDecreasing = append(safetyOrderDecreasing, currentItem)
				safetyCount++
			} else {
				unsafeOrderDecreasing = append(unsafeOrderDecreasing, currentItem)
			}

		}

		if len(safetyOrderDecreasing) == len(splitLine) {
			safeDecrease = append(safeDecrease, safetyOrderDecreasing)
		} else if len(unsafeOrderDecreasing) == len(splitLine) {
			unsafeDecrease = append(unsafeDecrease, unsafeOrderDecreasing)
		}
	}

	return safeDecrease, unsafeDecrease
}
